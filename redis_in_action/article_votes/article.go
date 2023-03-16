package article_votes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"strconv"
	"strings"
	"time"
)

const (
	OneWeekInSeconds = 7 * 3600 * 24
	VoteScore        = 432
)

var (
	idKey            = "articleId"
	scoreKey         = "score"
	timeKey          = "time"
	articleKeyPrefix = "article:"
	votedKeyPrefix   = "voted:"
	groupPrefix      = "group:"
)

type ArticleCmd interface {
	ArticleVote(articleKey, user string) error
	PostArticle(user, title, link string) (string, error)
	GetArticle(key string) (Article, error)
	GetArticles(page, size int64, order string) ([]Article, error)
	AddRemoveGroups(articleId string, toAdd, toRemove []string) error
	Reset() error
}

type Article struct {
	Id         int64
	Title      string
	Link       string
	Author     string
	CreateTime time.Time
	//Votes      int64
}

func (a Article) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a Article) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &a)
}

// ArticleRepo 文章仓库，其存储设计为：
// 一个按时间（time）排序的article列表
// 一个按分值（time+VoteScore）排序的article列表
// 有一个id生成的记录
// redis key设计：采用 [类别]:id 的形式，如 voted:1, article:1
type ArticleRepo struct {
	conn redis.Cmdable
}

var _ ArticleCmd = &ArticleRepo{}

// ArticleVote 给文章投票
// 投票逻辑为：
// 1. 判断文章是否为1周前发表
// 2. 检查当前用户是否投过票
// 3. 投票+1
func (a *ArticleRepo) ArticleVote(articleKey, user string) error {
	err := a.checkArticleVotable(articleKey)
	if err != nil {
		return err
	}

	articleId := strings.Split(articleKey, ":")[1]
	err = a.checkArticleRepeatableVote(articleId, user)
	if err != nil {
		return err
	}

	return a.vote(articleId)
}
func (a *ArticleRepo) checkArticleVotable(articleKey string) error {
	postedAt, err := a.conn.ZScore(context.Background(), timeKey, articleKey).Result()
	if err != nil {
		return err
	}
	cutoff := time.Now().Unix() - OneWeekInSeconds
	if postedAt < float64(cutoff) {
		return errors.New("article is posted one week ago")
	}
	return nil
}

func (a *ArticleRepo) checkArticleRepeatableVote(articleId, user string) error {
	votedKey := votedKeyPrefix + articleId
	res, err := a.conn.SAdd(context.Background(), votedKey, user).Result()
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("voted repeatably")
	}
	return nil
}

func (a *ArticleRepo) vote(articleId string) error {
	articleKey := articleKeyPrefix + articleId
	err := a.conn.ZIncrBy(context.Background(), scoreKey, VoteScore, articleKey).Err()
	if err != nil {
		return err
	}
	return a.conn.IncrBy(context.Background(), articleKey, 1).Err()
}

func (a *ArticleRepo) PostArticle(user, title, link string) (string, error) {
	id, err := a.conn.Incr(context.Background(), idKey).Result()
	if err != nil {
		return "", err
	}
	articleId := strconv.FormatInt(id, 10)
	now := time.Now()
	article := Article{
		Id:         id,
		Title:      title,
		Link:       link,
		Author:     user,
		CreateTime: now,
		//Votes:      1,
	}
	articleKey := articleKeyPrefix + articleId
	votedKey := votedKeyPrefix + articleId
	a.conn.SAdd(context.Background(), votedKey, user)
	a.conn.Expire(context.Background(), votedKey, time.Second*OneWeekInSeconds)

	a.conn.Set(context.Background(), articleKey, fmt.Sprintf("%v", article), redis.KeepTTL)
	_, err = a.conn.Set(context.Background(), articleKey, article, redis.KeepTTL).Result()
	if err != nil {
		return articleId, err
	}

	_, err = a.conn.ZAdd(context.Background(), scoreKey, redis.Z{
		Score:  float64(now.Unix() + VoteScore),
		Member: articleKey,
	}).Result()
	if err != nil {
		return articleId, err
	}

	_, err = a.conn.ZAdd(context.Background(), timeKey, redis.Z{
		Score:  float64(now.Unix()),
		Member: articleKey,
	}).Result()
	if err != nil {
		return articleId, err
	}
	return articleId, nil
}

func (a *ArticleRepo) GetArticle(key string) (Article, error) {
	article := Article{}
	err := a.conn.Get(context.Background(), key).Scan(&article)
	return article, err
}

func (a *ArticleRepo) GetArticles(page, size int64, order string) ([]Article, error) {
	if order != "score" {
		order = "time"
	}
	start := (page - 1) * size
	end := page*size - 1
	articleKeys, err := a.conn.ZRange(context.Background(), order, start, end).Result()
	if err != nil {
		return nil, err
	}
	articles := make([]Article, 0, len(articleKeys))
	//for _, articleKey := range articleKeys {
	//	article := Article{}
	//	err = a.conn.Get(context.Background(), articleKey).Scan(&article)
	//	if err != nil {
	//		return nil, err
	//	}
	//	articles = append(articles, article)
	//}

	res, err := a.conn.MGet(context.Background(), articleKeys...).Result()
	for _, r := range res {
		d, ok := r.(Article)
		if !ok {
			return nil, errors.New("store error")
		}
		articles = append(articles, d)
	}

	return articles, nil

}

func (a *ArticleRepo) AddRemoveGroups(articleId string, toAdd, toRemove []string) error {
	articleKey := articleKeyPrefix + articleId
	err := addRemoveProxy(toAdd, articleKey, a.conn.SAdd)
	if err != nil {
		return err
	}

	return addRemoveProxy(toRemove, articleKey, a.conn.SRem)
}

func addRemoveProxy(list []string, key string, f func(ctx context.Context, key string, members ...any) *redis.IntCmd) error {
	ctx := context.Background()
	for _, group := range list {
		groupKey := groupPrefix + group
		err := f(ctx, groupKey, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *ArticleRepo) Reset() error {
	return a.conn.FlushDB(context.Background()).Err()
}
