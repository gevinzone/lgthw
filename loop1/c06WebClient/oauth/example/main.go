package main

import (
	"context"
	"github.com/gevinzone/lgthw/loop1/c06WebClient/oauth"
)

func main() {
	conf := oauth.Setup()
	token, err := oauth.GetToken(context.Background(), conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(context.Background(), token)
	if err = oauth.GetUser(client); err != nil {
		panic(err)
	}
}
