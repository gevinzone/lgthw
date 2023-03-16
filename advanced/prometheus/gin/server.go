package ginprome

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func StartServer(engine *gin.Engine, addr string) error {
	RegisterMetrics()
	engine.Use(promMiddleware())
	bs(engine)
	metrics(engine)
	return engine.Run(addr)

}

func startMetrics(addr string) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addr, nil)
}
