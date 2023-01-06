package basic

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_ID"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Endpoint:     github.Endpoint,
		Scopes:       []string{"repo", "user"},
	}
}

func GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	url := conf.AuthCodeURL("state1")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
