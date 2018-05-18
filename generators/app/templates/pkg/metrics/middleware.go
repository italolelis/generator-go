package metrics

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"<%=projectRoot%>/pkg/log"
	"github.com/uber-go/tally"
)

// NewMiddleware creates a new log middleware
func NewMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.WithContext(r.Context())
		scope := WithContext(r.Context())

		logger.Debug("Metrics middleware started")
		h := scope.Histogram("requests", tally.DefaultBuckets)

		hsw := h.Start()
		m := httpsnoop.CaptureMetrics(handler, w, r)
		h.RecordDuration(m.Duration)
		hsw.Stop()

		logger.Debug("Metrics middleware finished")
	})
}
