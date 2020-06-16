package foundation

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

// InitTracer to initialize tracing API
func InitTracer(config ConfigStore) {
	initJaegerTracer(config)
}

func initJaegerTracer(config ConfigStore) {
	serviceName := config.GetConfig("app.name").(string)
	enableTracing := config.GetConfig("tracer.url").(bool)
	if enableTracing {
		tracerURL := config.GetConfig("tracer.url").(string)
		// init open tracing
		udpTransport, _ := jaeger.NewUDPTransport(tracerURL, 0)
		reporter := jaeger.NewRemoteReporter(udpTransport)
		sampler := jaeger.NewConstSampler(true)
		tracer, _ := jaeger.NewTracer(serviceName, sampler, reporter)
		opentracing.SetGlobalTracer(tracer)
	}
}

// IntroduceSpan to start new tracing context
func IntroduceSpan(ctx context.Context, spanName string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, spanName)
	return span, ctx
}

//Serialise to generate tracing context
func Serialise(ctx context.Context, req *http.Request) {
	req = req.WithContext(ctx)
	if span := opentracing.SpanFromContext(ctx); span != nil {
		opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(req.Header))
	}
}

//Deserialize to process tracing context
func Deserialize(r *http.Request, spanName string) (opentracing.Span, *http.Request) {
	time.Sleep(250 * time.Millisecond)
	var serverSpan opentracing.Span
	appSpecificOperationName := spanName
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		log.Fatal("Error")
	}
	serverSpan = opentracing.StartSpan(
		appSpecificOperationName,
		ext.RPCServerOption(wireContext))
	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	r = r.WithContext(ctx)
	return serverSpan, r
}

// zipkin "github.com/openzipkin/zipkin-go"
// zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

/*
// InitTracer to initialize tracing API
func InitTracer(config ConfigStore) {
	zipkinURL := config.GetConfig("zipkin-url").(string)
	collector, err := zipkin.NewHTTPCollector(fmt.Sprintf("%s/api/v1/spans", zipkinURL))
	if err != nil {
		panic("Error connecting to zipkin server at " +
			fmt.Sprintf("%s/api/v1/spans", zipkinURL) + ". Error: " + err.Error())
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "127.0.0.1:0", serviceName))
	if err != nil {
		panic("Error starting new zipkin tracer. Error: " + err.Error())
	}
}

// TracingMiddleware Used for request tracing purposes
func TracingMiddleware(tracer *zipkin.Tracer) http.Handler {
	return zipkinhttp.NewServerMiddleware(tracer, zipkinhttp.SpanName("request"))
}
*/
