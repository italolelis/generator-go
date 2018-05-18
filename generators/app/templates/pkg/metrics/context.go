package metrics

import (
	"context"
	"io"
	"time"

	"github.com/go-chi/chi"

	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
)

type metricsKeyType int

const metricsKey metricsKeyType = iota

var reporter prometheus.Reporter

func init() {
	reporter = prometheus.NewReporter(prometheus.Options{})
}

// NewContext creates a new metrics context
func NewContext(ctx context.Context, prefix string) (context.Context, io.Closer) {
	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Prefix:         prefix,
		CachedReporter: reporter,
		Separator:      prometheus.DefaultSeparator,
	}, 1*time.Second)

	return context.WithValue(ctx, metricsKey, scope), closer
}

// WithContext gets the existing metrics from context
func WithContext(ctx context.Context) tally.Scope {
	ctxMetrics, _ := ctx.Value(metricsKey).(tally.Scope)
	return ctxMetrics
}

// RegisterRoutes registers prometheus routes
func RegisterRoutes(r chi.Router) {
	r.Handle("/metrics", reporter.HTTPHandler())
}
