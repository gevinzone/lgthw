package ginprome

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"strconv"
	"time"
)

func promMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		mockDuration := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(mockDuration))
		c.Next()
		defer func() {
			res := time.Now().Sub(start).Microseconds()
			// Counter
			AccessCounter.With(prometheus.Labels{
				"method": c.Request.Method,
				"path":   c.Request.URL.Path,
			}).Add(1)

			// Histogram
			HttpDurationsHistogram.With(prometheus.Labels{
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
				"status": strconv.Itoa(c.Writer.Status()),
			}).Observe(float64(res))

			// Summary
			HttpDurations.With(prometheus.Labels{
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
				"status": strconv.Itoa(c.Writer.Status()),
			}).Observe(float64(res))
		}()

	}
}
