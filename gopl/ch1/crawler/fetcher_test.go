package crawler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	testcases := []struct {
		name    string
		url     string
		wantErr error
	}{
		{
			name: "normal",
			url:  "https://blog.igevin.info",
		},
		{
			name:    "invalid",
			url:     "https://igevin.info",
			wantErr: errors.New("fail to fetch"),
		},
	}
	for _, tc := range testcases {
		b, err := fetch(tc.url)
		assert.Equal(t, tc.wantErr, err)
		if err != nil {
			return
		}
		assert.NotNil(t, b)
	}
}

func TestFetchAll(t *testing.T) {
	urls := []string{
		"https://www.baidu.com",
		"https://igevin.info",
		"https://blog.igevin.info",
		"https://douban.com",
	}
	ch := make(chan string, len(urls))
	fetchAll(urls, ch)
	time.Sleep(time.Second)
	assert.Equal(t, len(ch), len(urls))
}
