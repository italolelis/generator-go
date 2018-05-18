package tracer

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"<%=projectRoot%>/pkg/log"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// NewTracerRequest is the middleware function for OpenTracing
func NewTracerRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var span opentracing.Span
		var err error

		logger := log.WithContext(r.Context())
		logger.Debug("Tracer middleware started")
		// Attempt to join a trace by getting trace context from the headers.
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			// If for whatever reason we can't join, go ahead an start a new root span.
			span = opentracing.StartSpan(r.URL.Path)
		} else {
			span = opentracing.StartSpan(r.URL.Path, opentracing.ChildOf(wireContext))
		}
		defer span.Finish()

		host, err := os.Hostname()
		if err != nil {
			logger.With(err).Warn("Failed to get host name")
		}

		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(
			span,
			fmt.Sprintf("%s://%s%s", r.URL.Scheme, r.Host, r.URL.Path),
		)
		ext.Component.Set(span, "test")
		ext.SpanKind.Set(span, "server")

		span.SetTag("peer.address", r.RemoteAddr)
		span.SetTag("host.name", host)

		// Add information on the peer service we're about to contact.
		if host, portString, err := net.SplitHostPort(r.URL.Host); err == nil {
			ext.PeerHostname.Set(span, host)
			if port, err := strconv.Atoi(portString); err != nil {
				ext.PeerPort.Set(span, uint16(port))
			}
		} else {
			ext.PeerHostname.Set(span, r.URL.Host)
		}

		err = span.Tracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			logger.With(err).Error("Could not inject span context into header")
		}

		handler.ServeHTTP(w, r.WithContext(opentracing.ContextWithSpan(r.Context(), span)))
		logger.Debug("Tracer middleware finished")
	})
}
