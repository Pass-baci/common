package common

import (
	"github.com/asim/go-micro/v3/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

type JaegerConfig struct {
	ServiceName string `json:"service_name"`
	Address     string `json:"address"`
}

func NewTracer(serviceName, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
		ServiceName: serviceName,
	}

	return cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
}

// GetJaegerConfigFromConsul 获取jaeger配置
func GetJaegerConfigFromConsul(config config.Config, path ...string) (*JaegerConfig, error) {
	jaegerConfig := &JaegerConfig{}
	if err := config.Get(path...).Scan(jaegerConfig); err != nil {
		return nil, err
	}
	return jaegerConfig, nil
}
