package storage

import (
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"os"
	"sync"
)

type FileStorage struct {
	Path string
	mu   sync.RWMutex
}

func (f *FileStorage) GetToken() (*oauth2.Token, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	in, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	var t *oauth2.Token
	decoder := json.NewDecoder(in)
	//err = decoder.Decode(t)
	//return t, err
	return t, decoder.Decode(t)
}

func (f *FileStorage) SetToken(t *oauth2.Token) error {
	if t == nil || !t.Valid() {
		return errors.New("bad token")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	out, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()
	data, err := json.Marshal(&t)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	return err
}
