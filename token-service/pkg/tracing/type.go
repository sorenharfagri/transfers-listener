package tracing

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"token-service/pkg/type/context"
	log "token-service/pkg/type/logger"
)

type Options struct {
	JaegerHost  string
	JaegerPort  uint32
	ServiceName string
}

func New(ctx context.Context, options Options) (io.Closer, error) {

	cfg := &config.Configuration{
		ServiceName: options.ServiceName,
		RPCMetrics:  true,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", options.JaegerHost, options.JaegerPort),
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}
