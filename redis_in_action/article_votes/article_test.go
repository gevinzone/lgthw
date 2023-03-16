package article_votes

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestArticleRepo_PostArticle(t *testing.T) {
	client := redisClient(t)
	repo := ArticleRepo{conn: client}
	user, title, link := "user", "title", "link"
	articleId, err := repo.PostArticle(user, title, link)
	assert.NoError(t, err)

	votedKey := votedKeyPrefix + articleId
	articleKey := articleKeyPrefix + articleId
	assert.Equal(t, int64(0), client.SAdd(context.Background(), votedKey, user).Val())
	res, err := client.ZScore(context.Background(), scoreKey, articleKey).Result()
	require.NoError(t, err)
	require.Less(t, float64(0), res)
	require.NoError(t, client.ZScore(context.Background(), timeKey, articleKey).Err())
}

func TestArticleRepo_GetArticle(t *testing.T) {
	client := redisClient(t)
	repo := ArticleRepo{conn: client}
	user, title, link := "user0", "title0", "link0"
	id, err := repo.PostArticle(user, title, link)
	require.NoError(t, err)
	t.Log(id)
	a, err := repo.GetArticle(id)
	require.NoError(t, err)
	assert.Equal(t, id, strconv.FormatInt(a.Id, 10))
	assert.Equal(t, title, a.Title)
	//t.Log(a)
}

func TestArticleRepo_Reset(t *testing.T) {
	client := redisClient(t)
	repo := ArticleRepo{conn: client}
	err := repo.Reset()
	require.NoError(t, err)
}

func redisClient(t *testing.T) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	_, err := client.Ping(context.Background()).Result()
	require.NoError(t, err)
	return client
}
