//go:build linux

package main

import (
	ginprome "github.com/gevinzone/lgthw/advanced/prometheus/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	err := ginprome.StartServer(engine, ":8080")
	if err != nil {
		panic(err)
	}
}
