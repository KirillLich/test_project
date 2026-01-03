package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	},
	[]string{"code", "method", "path"},
)

type ResponseWriterInterceptor struct {
	http.ResponseWriter
	StatusCode int
}

func PrometeusMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &ResponseWriterInterceptor{ResponseWriter: w, StatusCode: http.StatusOK}
		handler.ServeHTTP(ww, r)

		requestsTotal.With(
			prometheus.Labels{
				"code":   strconv.Itoa(ww.StatusCode),
				"method": r.Method,
				"path":   r.URL.Path,
			}).Inc()
	})
}

func main() {
	prometheus.MustRegister(requestsTotal)

	fmt.Println("Listen and serve: localhost:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello, SRE!") })
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})
	mux.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", PrometeusMiddleware(mux))
}
