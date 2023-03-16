package ginprome

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	err := StartServer(engine, ":8080")
	assert.NoError(t, err)
}

//func startServer(engine *gin.Engine) error {
//	RegisterMetrics()
//	engine.Use(promMiddleware())
//	// 业务
//	bs(engine)
//	// metrics
//	metrics(engine)
//
//	return engine.Run()
//
//}
