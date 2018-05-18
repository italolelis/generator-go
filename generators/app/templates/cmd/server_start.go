package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gojektech/heimdall"
	"github.com/hashicorp/errwrap"
	"<%=projectRoot%>/pkg/config"
	"<%=projectRoot%>/pkg/handler"
	"<%=projectRoot%>/pkg/log"
	"<%=projectRoot%>/pkg/tracer"
	"<%=projectRoot%>/pkg/metrics"
	"github.com/spf13/cobra"
)

type (
	// StartServerOpts are the flags for the start server command
	StartServerOpts struct{}
)

// NewStartServerCmd starts the web server
func NewStartServerCmd(ctx context.Context) *cobra.Command {
	opts := StartServerOpts{}

	cmd := cobra.Command{
		Use:   "start",
		Short: "Starts a web server for captain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts)
		},
	}

	return &cmd
}

// RunServerStart starts a web server command
func RunServerStart(ctx context.Context, opts StartServerOpts) error {
	logger := log.WithContext(ctx)
	defer logger.Sync()
	cfg := config.WithContext(ctx)

	err := tracer.NewContext(ctx)
	if err != nil {
		return errwrap.Wrap(err, errors.New("failed to init tarcer"))
	}

	ctx, closer := metrics.NewContext(ctx, "")
	defer closer.Close()

	httpClient := initHTTPClient(ctx)

	logger.Info("Starting service")
	// creates a web router
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		tracer.NewTracerRequest,
		metrics.NewMiddleware,
		log.NewMiddleware,
	)
	metrics.RegisterRoutes(r)

	r.Get("/docs", handler.Docs)
	r.Get("/hystrix", buildHystrixHandler())
	// r.Get("/status", health.HandlerFunc)
	r.Get("/", handler.HelloWorld)
	r.Get("/github/repos", handler.GithubRepos(httpClient))

	logger.Infow("Service running", "port", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), chi.ServerBaseContext(ctx, r))
}

func initHTTPClient(ctx context.Context) heimdall.Client {
	logger := log.WithContext(ctx)
	cfg := config.WithContext(ctx)

	fallbackFn := func(err error) error {
		logger.With(err).Warn("Circuit triped")
		return err
	}

	hystrixConfig := heimdall.NewHystrixConfig("github", heimdall.HystrixCommandConfig{
		Timeout:                cfg.CircuitBreaker.Timeout,
		MaxConcurrentRequests:  cfg.CircuitBreaker.MaxConcurrentRequests,
		ErrorPercentThreshold:  cfg.CircuitBreaker.ErrorPercentThreshold,
		SleepWindow:            cfg.CircuitBreaker.SleepWindow,
		RequestVolumeThreshold: cfg.CircuitBreaker.RequestVolumeThreshold,
		FallbackFunc:           fallbackFn,
	})

	timeout := 1000 * time.Millisecond
	client := heimdall.NewHystrixHTTPClient(timeout, hystrixConfig)
	client.SetRetrier(buildRetrier(cfg.Retry))
	client.SetRetryCount(cfg.Retry.Attempts)

	return client
}

func buildRetrier(cfg config.Retry) heimdall.Retriable {
	backoff := heimdall.NewExponentialBackoff(
		cfg.InitialTimeout,
		cfg.MaxTimeout,
		cfg.ExponentFactory,
		cfg.MinimumJitterInterval,
	)

	// Create a new retry mechanism with the backoff
	return heimdall.NewRetrier(backoff)
}

func buildHystrixHandler() http.HandlerFunc {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	return hystrixStreamHandler.ServeHTTP
}
