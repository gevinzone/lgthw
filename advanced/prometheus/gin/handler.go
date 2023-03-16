package ginprome

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func metrics(r gin.IRouter) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func bs(r gin.IRouter) {
	r.GET("/test/get/*get", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.POST("/test/post/*post", func(c *gin.Context) {
		c.String(http.StatusCreated, "OK")
	})
	r.PUT("/test/put/*put", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.DELETE("/test/delete/*delete", func(c *gin.Context) {
		c.String(http.StatusNoContent, "OK")
	})
}
