package rest

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/avinash92c/bootstrap-go/model"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterProfRoutes registers all memory profiling routes for http profiling
func RegisterProfRoutes(router *model.Router) {
	/*
		router.Router.HandleFunc("/debug/pprof/", pprof.Index)
		router.Router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.Router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.Router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

		router.Router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		router.Router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		router.Router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		router.Router.Handle("/debug/pprof/block", pprof.Handler("block"))
	*/

	router.Router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
}

// RegisterTelemetryRoutes contains metrics & telemetry routes
func RegisterTelemetryRoutes(router *model.Router) {
	router.Router.Handle("/metrics", promhttp.Handler())
	router.Router.Path("/health").HandlerFunc(healthCheck).Methods("GET")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"UP"}`)
}
