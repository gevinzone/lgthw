package crawler

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("fail to fetch")
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func fetchAll(urls []string, ch chan string) {
	for _, url := range urls {
		go func(url string) {
			b, err := fetch(url)
			if err != nil {
				ch <- err.Error()
				return
			}
			ch <- string(b)
		}(url)
	}
}
