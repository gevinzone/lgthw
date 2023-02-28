package basic

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// 可以直接用counter
	opsProceeded = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	// 也可以通过 Vector 创建counter
	accessCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "access_counter",
		},
		[]string{"method", "path"},
	)
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	opsProceeded.Inc()
	_, _ = w.Write([]byte("hello, world"))
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	accessCounterVec.With(prometheus.Labels{
		"method": r.Method,
		"path":   r.URL.Path,
	}).Inc()
	_, _ = w.Write([]byte("hello from other pages"))
}

func registerPrometheus() {
	prometheus.MustRegister(accessCounterVec)
}

func ListenAndServe(addr string) error {
	http.HandleFunc("/root", handleRoot)

	// 使用vector时，要先注册了才能用
	registerPrometheus()
	http.HandleFunc("/other", handleOther)

	// 不要把prometheus 的监控数据暴露出去
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":8090", nil)
	}()

	return http.ListenAndServe(addr, nil)
}
