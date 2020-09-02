package tracing_wrapper

import (
	"bytes"
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"reflect"
)

var closer io.Closer

func InitJaeger(service string) {

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
			//LocalAgentHostPort: "192.168.87.159:6831",
		},
		ServiceName: service,
	}
	var tracer opentracing.Tracer
	var err error
	tracer, closer, err = cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
}

func TracingRootSpan(val reflect.Value, spanName string, key string, value string) (opentracing.Span, error) {
	var span opentracing.Span = nil
	var err error = nil

	span = opentracing.StartSpan(spanName)
	if val.CanSet() {
		buf := bytes.NewBufferString("")
		err = opentracing.GlobalTracer().Inject(span.Context(), opentracing.Binary, buf)
		if err != nil {
			return nil, err
		}
		val.SetBytes(buf.Bytes())
		span.SetTag(key, value)
	} else {
		fmt.Println("TracingRootSpan val.CanSet().else")
	}

	return span, err
}

func TracingChildSpan(val reflect.Value, spanName string, key string, value string) (opentracing.Span, error) {
	var span opentracing.Span = nil
	var err error = nil
	if val.CanSet() {
		ctx, err := opentracing.GlobalTracer().Extract(opentracing.Binary, bytes.NewBuffer(val.Bytes()))
		if err == nil {
			span = opentracing.StartSpan(spanName, ext.RPCServerOption(ctx))
			span.SetTag(key, value)

			//将新的span数据写入msg
			buf := bytes.NewBuffer([]byte{})
			err = opentracing.GlobalTracer().Inject(span.Context(), opentracing.Binary, buf)
			if err != nil {
				val.SetBytes(buf.Bytes())
			}
		}
	} else {
		fmt.Println("TracingChildSpan val.CanSet().else")
	}
	return span, err
}

func TracingEndSpan(span opentracing.Span) {
	if span != nil {
		span.Finish()
	}
}

func CloseJaeger() {

	closer.Close()
}
