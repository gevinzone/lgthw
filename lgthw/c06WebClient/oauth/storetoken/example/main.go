package main

import (
	"context"
	"github.com/gevinzone/lgthw/loop1/c06WebClient/oauth/storetoken"
	"github.com/gevinzone/lgthw/loop1/c06WebClient/oauth/storetoken/storage"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"os"
)

func main() {
	conf := storetoken.Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_ID"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Endpoint:     github.Endpoint,
			Scopes:       []string{"repo", "user"},
		},
		Storage: &storage.FileStorage{Path: "./loop1/c06WebClient/oauth/storetoken/storage/token.txt"},
	}
	ctx := context.Background()
	token, err := storetoken.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
