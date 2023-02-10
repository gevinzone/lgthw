package main

import (
	"context"
	"github.com/gevinzone/lgthw/lgthw/c06WebClient/oauth/basic"
)

func main() {
	conf := basic.Setup()
	token, err := basic.GetToken(context.Background(), conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(context.Background(), token)
	if err = basic.GetUser(client); err != nil {
		panic(err)
	}
}
