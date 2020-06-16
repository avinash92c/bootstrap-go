package model

import (
	"github.com/avinash92c/bootstrap-go/database"
	"github.com/avinash92c/bootstrap-go/foundation"
	"github.com/gorilla/mux"
)

//AppServer contains all app level common dependencies
type AppServer struct {
	Db          *database.DB
	Config      foundation.ConfigStore
	Logger      foundation.Logger
	ServiceName string
}

// NewAppServer instance
func NewAppServer(serviceName string, db *database.DB, config foundation.ConfigStore, logger foundation.Logger) *AppServer {
	return &AppServer{ServiceName: serviceName, Db: db, Config: config, Logger: logger}
}

//Router dedicated type to reduce code refactor for router swap
type Router struct { //NOT THE BEST WAY, BUT USE FOR NOW
	Router    *mux.Router
	AppRouter *mux.Router
}

//NewRouter instance
func NewRouter() *Router {
	return &Router{Router: mux.NewRouter()}
}

//MakeAppRouter Generates a subrouter for application specific endpoints
func (router *Router) MakeAppRouter(path string) {
	router.AppRouter = router.Router.PathPrefix(path).Subrouter()
}

//SubRouter Generate a New Subrouter
func (router *Router) SubRouter(path string) *Router {
	return &Router{Router: router.Router.PathPrefix(path).Subrouter()}
}
