package storetoken

import (
	"context"
	"golang.org/x/oauth2"
)

type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// Token satisfies the TokenSource interface
func (s *storageTokenSource) Token() (*oauth2.Token, error) {
	var (
		token *oauth2.Token
		err   error
	)
	if token, err = s.GetToken(); err == nil && token.Valid() {
		return token, nil
	}
	if token, err = s.TokenSource.Token(); err != nil {
		return nil, err
	}
	if err = s.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// StorageTokenSource will be used by our config.TokenSource method
func StorageTokenSource(ctx context.Context, c *Config, t *oauth2.Token) oauth2.TokenSource {
	if t == nil || !t.Valid() {
		if token, err := c.GetToken(); err == nil {
			t = token
		}
	}
	ts := c.Config.TokenSource(ctx, t)
	return &storageTokenSource{
		Config:      c,
		TokenSource: ts,
	}
}
