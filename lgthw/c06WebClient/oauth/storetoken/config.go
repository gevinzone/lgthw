package storetoken

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

type Config struct {
	*oauth2.Config
	Storage
}

func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err = c.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// TokenSource can be passed a token which
// is stored, or when a new one is retrieved,
// that's stored
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, t)
}

func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}

type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(token *oauth2.Token) error
}

func GetToken(ctx context.Context, conf Config) (*oauth2.Token, error) {
	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, nil
	}
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")
	var code string
	if _, err = fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
