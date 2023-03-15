package article_votes

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"strconv"
	"strings"
	"time"
)

const (
	OneWeekInSeconds       = 7 * 3600 * 24
	VoteScore              = 432
	ArticlesPerPage  int64 = 25
)

var (
	idKey            = "article:"
	scoreKey         = "score"
	timeKey          = "time"
	articleKeyPrefix = "article:"
	votedKeyPrefix   = "voted:"
)

type ArticleCmd interface {
	ArticleVote(articleKey, user string) error
	PostArticle(user, title, link string) (string, error)
	GetArticles(page, size int64, order string) ([]Article, error)
	AddRemoveGroups(articleId string, toAdd, toRemove []string)
	GetGroupArticles(group, order string, page int64) ([]Article, error)
	Reset()
}

type Article struct {
	Id         int64
	Title      string
	Link       string
	Author     string
	CreateTime time.Time
	Votes      int64
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
	postedAt, err := a.conn.ZScore(context.Background(), timeKey, articleKey).Result()
	if err != nil {
		return err
	}
	cutoff := time.Now().Unix() - OneWeekInSeconds
	if postedAt < float64(cutoff) {
		return errors.New("article is posted one week ago")
	}
	articleId := strings.Split(articleKey, ":")[1]
	votedKey := votedKeyPrefix + articleId
	var res int64
	if res, err = a.conn.SAdd(context.Background(), votedKey, user).Result(); err != nil {
		return err
	}
	if res == 0 {
		return errors.New("voted repeatably")
	}
	a.conn.ZIncrBy(context.Background(), scoreKey, VoteScore, articleKey)
	a.conn.IncrBy(context.Background(), articleKey, 1)
	return nil
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
		Votes:      1,
	}
	articleKey := articleKeyPrefix + articleId
	votedKey := votedKeyPrefix + articleId
	a.conn.SAdd(context.Background(), votedKey, user)
	a.conn.Expire(context.Background(), votedKey, time.Second*OneWeekInSeconds)

	a.conn.Set(context.Background(), articleKey, article, redis.KeepTTL)

	a.conn.ZAdd(context.Background(), scoreKey, redis.Z{
		Score:  float64(now.Unix() + VoteScore),
		Member: articleKey,
	})
	a.conn.ZAdd(context.Background(), timeKey, redis.Z{
		Score:  float64(now.Unix()),
		Member: articleKey,
	})
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

func (a *ArticleRepo) AddRemoveGroups(articleId string, toAdd, toRemove []string) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) GetGroupArticles(group, order string, page int64) ([]Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) Reset() {
	a.conn.FlushDB(context.Background())
}
