package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of latencies for HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

// MetricsMiddleware instruments requests with Prometheus metrics.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start).Seconds()
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		// Avoid instrumenting the metrics endpoint itself and common static/swagger paths
		// Check both route template (path) and raw URL path to cover misspellings like /metrucs
		rawPath := c.Request.URL.Path
		if path == "/metrics" || strings.HasPrefix(path, "/swagger") || path == "/favicon.ico" ||
			rawPath == "/metrics" || strings.HasPrefix(rawPath, "/swagger") || rawPath == "/favicon.ico" {
			return
		}

		// Normalize empty or unmatched routes to reduce label cardinality
		if path == "" {
			path = "unmatched"
		}

		status := c.Writer.Status()
		statusStr := strconv.Itoa(status)

		httpRequestsTotal.WithLabelValues(method, path, statusStr).Inc()
		httpRequestDuration.WithLabelValues(method, path).Observe(latency)
	}
}
