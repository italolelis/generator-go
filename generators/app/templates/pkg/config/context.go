package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type configKeyType int

const configKey configKeyType = iota

// NewContext loads a configuration file into the Spec struct and return a context
func NewContext(ctx context.Context, configFile string) (context.Context, error) {
	var config Spec

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, configKey, &config), nil
}

// WithContext returns a logrus logger from the context
func WithContext(ctx context.Context) *Spec {
	if ctx == nil {
		return nil
	}

	if ctxConfig, ok := ctx.Value(configKey).(*Spec); ok {
		return ctxConfig
	}

	return nil
}
