package rest

import (
	"net/http"

	"github.com/avinash92c/bootstrap-go/model"
	"github.com/avinash92c/bootstrap-go/security"
	"github.com/justinas/alice"
	"github.com/opentracing/opentracing-go"
)

//RoutingEngine interface defines interfaces for configuring rest routes
type RoutingEngine interface {
	BuildRoute(path, method string, secure bool, handler Handler)
	GenerateRouteHandler(handler Handler, secure bool) http.Handler
	TracingMiddleware(next http.Handler) http.Handler
}

// NewRoutingEngine returns a pointer to RoutingEngine Interface
func NewRoutingEngine(router *model.Router, appserver *model.AppServer) RoutingEngine {
	return &routingengine{appserver: appserver, router: router}
}

type routingengine struct {
	router    *model.Router
	appserver *model.AppServer
}

func (routeEngine *routingengine) BuildRoute(path, method string, secure bool, handler Handler) {
	routeEngine.router.AppRouter.Path(path).Handler(routeEngine.GenerateRouteHandler(handler, secure)).Methods(method)
}

//generateRoute configures your route with requested middleware and some standard middleware
func (routeEngine *routingengine) GenerateRouteHandler(handler Handler, secure bool) http.Handler {
	routechain := alice.New()
	appserver := routeEngine.appserver
	// routechain = routechain.Append(appserver.TracingMiddleware)
	if secure { //FOR NOW TO CHECK
		routechain = routechain.Append(security.TokenCheck)
	}
	return routechain.ThenFunc(generateHandler(appserver, handler))
}

//Handler Interface to Inject Handler Functions
type Handler func(*model.AppServer, http.ResponseWriter, *http.Request)

//Handle Function Segment
func (f Handler) Handle(appserver *model.AppServer, w http.ResponseWriter, r *http.Request) {
	f(appserver, w, r)
}

func generateHandler(appserver *model.AppServer, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// serverSpan, r := foundation.Deserialize(r, appserver.ServiceName) //TODO LATER
		// defer serverSpan.Finish()

		handler.Handle(appserver, w, r)
	}
}

// TracingMiddleware Used for request tracing purposes
func (routeEngine *routingengine) TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appserver := routeEngine.appserver
		span, newCtx := opentracing.StartSpanFromContext(r.Context(), appserver.ServiceName)
		defer func() {
			span.Finish()
		}()
		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}
