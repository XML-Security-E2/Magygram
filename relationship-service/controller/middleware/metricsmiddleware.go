package middleware

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"sync"
)

type MetricsMiddleware struct {
	requestsProcessed *prometheus.CounterVec
	mutex        sync.RWMutex
}

func NewMetricsMiddleware() *MetricsMiddleware {
	requestsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "relationship_service_requests_processed",
		Help: "The total number of processed requests",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		requestsProcessed: requestsProcessed,
	}
}

func (m *MetricsMiddleware) Metrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		m.mutex.Lock()
		defer m.mutex.Unlock()
		m.requestsProcessed.With(prometheus.Labels{"method": c.Request().Method, "path": c.Path(), "statuscode": strconv.Itoa(c.Response().Status)}).Inc()
		return nil
	}
}