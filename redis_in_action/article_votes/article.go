package article_votes

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strconv"
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
	ArticleVote(article, user string)
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
// 1.
//
func (a *ArticleRepo) ArticleVote(article, user string) {
	//TODO implement me
	panic("implement me")
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

	a.conn.HSet(context.Background(), articleKey, article)

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

func (a *ArticleRepo) GetArticles(page, size int64, order string) ([]Article, error) {
	panic("implement me")
	//if order != "score" {
	//	order = "time"
	//}
	//start := (page - 1) * size
	//end := page*size - 1
	//articleKeys, err := a.conn.ZRange(context.Background(), order, start, end).Result()
	//if err != nil {
	//	return nil, err
	//}
	//articles := make([]Article, 0, len(articleKeys))
	//for _, articleKey := range articleKeys {
	//	a.conn.HGet()
	//	articles = append(articles)
	//}

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
