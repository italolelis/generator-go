package tracer

import "go.uber.org/zap"

type jaegerLoggerAdapter struct {
	log *zap.SugaredLogger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.log.Error(msg)
}

// Infof adapts infof messages to logrus
func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.log.Infof(msg, args)
}
