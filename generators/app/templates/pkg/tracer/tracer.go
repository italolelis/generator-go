package tracer

import (
	"context"

	"<%=projectRoot%>/pkg/log"
	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

const (
	componentName = "captain"
)

// NewContext creates a new tracer
func NewContext(ctx context.Context) error {
	logger := log.WithContext(ctx)
	logger.Debug("Using Jaeger as tracing system")

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return err
	}

	t, _, err := cfg.New(
		componentName,
		jaegercfg.Logger(jaegerLoggerAdapter{logger}),
		jaegercfg.Metrics(metrics.NullFactory),
	)

	opentracing.InitGlobalTracer(t)

	return err
}
