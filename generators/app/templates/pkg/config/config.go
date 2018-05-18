package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	// Spec represents the global app configs
	Spec struct {
		Port           int             `envconfig:"PORT"`
		LogLevel       zap.AtomicLevel `envconfig:"LOG_LEVEL"`
		Metrics        Metrics
		CircuitBreaker CircuitBreaker
		Retry          Retry
	}

	// Metrics represents the metrics configs
	Metrics struct {
		Prefix string `envconfig:"METRICS_PREFIX"`
	}

	// CircuitBreaker holds the CB configs
	CircuitBreaker struct {
		Timeout                int `envconfig:"CB_TIMEOUT"`
		MaxConcurrentRequests  int `envconfig:"CB_MAX_CONCURRENT_REQUEST"`
		ErrorPercentThreshold  int `envconfig:"CB_ERROR_PERCENT_THRESHOLD"`
		SleepWindow            int `envconfig:"CB_SLEEP_WINDOW"`
		RequestVolumeThreshold int `envconfig:"CB_REQUEST_VOLUME_THRESHOLD"`
	}

	// Retry holds the CB retries configs
	Retry struct {
		Attempts              int           `envconfig:"RETRY_ATTEMPTS"`
		InitialTimeout        time.Duration `envconfig:"RETRY_INITIAL_TIMEOUT"`
		MaxTimeout            time.Duration `envconfig:"RETRY_MAX_TIMEOUT"`
		ExponentFactory       float64       `envconfig:"RETRY_EXPONENT_FACTORY"`
		MinimumJitterInterval time.Duration `envconfig:"RETRY_MINIMUM_JITTER_INTERVAL"`
	}
)

func init() {
	viper.SetDefault("port", 8080)
	viper.SetDefault("logLevel", zap.NewAtomicLevel())
	viper.SetDefault("circuitBreaker.Timeout", 1100)
	viper.SetDefault("circuitBreaker.MaxConcurrentRequests", 100)
	viper.SetDefault("circuitBreaker.ErrorPercentThreshold", 25)
	viper.SetDefault("circuitBreaker.SleepWindow", 10)
	viper.SetDefault("circuitBreaker.RequestVolumeThreshold", 10)

	viper.SetDefault("retry.InitialTimeout", "200ms")
	viper.SetDefault("retry.MaxTimeout", "600ms")
	viper.SetDefault("retry.ExponentFactory", 2.0)
	viper.SetDefault("retry.MinimumJitterInterval", "2ms")
}
