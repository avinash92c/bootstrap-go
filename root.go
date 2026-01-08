package bootstrap

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/avinash92c/bootstrap-go/database"
	"github.com/avinash92c/bootstrap-go/foundation"
	"github.com/avinash92c/bootstrap-go/model"
	"github.com/avinash92c/bootstrap-go/rest"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	//DB DRIVERS
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	serverport   = 8080
	writeTimeout = time.Second * 30
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
	logger       foundation.Logger
	config       foundation.ConfigStore

	// FOR GRACEFUL SHUTDOWN OF APP
	server     *http.Server
	db         *database.DB
	configpath *string
)

func parseFlags() {
	configpath = flag.String("configpath", "./config", "Path to folder containing config files")
	flag.Parse()
}

// Init For App Initializations
func Init(opts ...InitOption) (*model.AppServer, *model.Router) {
	parseFlags()

	cfg := &initConfig{}

	for _, opt := range opts {
		opt(cfg)
	}

	config, logger = foundation.Init(*configpath, cfg.logging)
	// foundation.InitTracer(config) //TODO MODIFY FOR ON\OFF FLAG
	db = database.GetConnectionPool(config)

	serviceName := config.GetConfig("app.name").(string)

	appsvr := model.NewAppServer(serviceName, db, config, logger)
	baserouter := model.NewRouter()

	rest.RegisterTelemetryRoutes(baserouter)
	rest.RegisterProfRoutes(baserouter)

	baserouter.MakeAppRouter("/app") //OVERRIDE BASE ROUTER WITH SUBROUTER

	handleSigterm(func() {
		//TODO CLEANUP OF RESOURCES
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err := database.ShutdownPool(db)
		logger.ErrorF(context.TODO(), "Error Occurred %v", err)

		logger.Info(context.TODO(), "shutting down")
		server.Shutdown(ctx)
	})

	return appsvr, baserouter
}

// Handles Ctrl+C or most other means of "controlled" shutdown gracefully. Invokes the supplied func before exiting.
func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

// StartServer Post PreConfigurations
func StartServer(appserver *model.AppServer, router *model.Router) {
	//GET CONFIGURED PORT FROM CONFIG //COMMAND LINE FLAG SUPPORT LATER
	//Priority Order
	//FLAG - 1 , Config - 2 , HardCoded -3

	cfgport := config.GetConfig("app.port").(int)
	if cfgport > 0 {
		serverport = cfgport
		logger.InfoF(context.TODO(), "Starting Server On Port %v", serverport)
	}

	// printRoutes(router.Router)

	server = &http.Server{
		Addr:         fmt.Sprintf(":%v", serverport),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      router.Router,
	}
	log.Fatal(server.ListenAndServe())
}

func startGRPCServer(appserver *model.AppServer) {
	logger.Info(context.TODO(), "Starting GRPC Server")
	cfgport := config.GetConfig("app.port").(int)
	if cfgport > 0 {
		serverport = cfgport
		logger.InfoF(context.TODO(), "Starting Server On Port %v", serverport)
	}
	svrport := fmt.Sprintf(":%v", serverport)
	listen, err := net.Listen("tcp", svrport)
	if err != nil {
		logger.Error(context.TODO(), err)
	}
	s := grpc.NewServer()
	// grpcservices.RegisterServices(s) //SERVICE REGISTRY
	if err := s.Serve(listen); err != nil {
		logger.Error(context.TODO(), "Failed To Start")
	}
	logger.Info(context.TODO(), "Started Server")
}

func printRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		logger.Info(context.TODO(), "ROUTE:", pathTemplate)
		methods, err := route.GetMethods()
		if err != nil {
			// return err
			logger.Error(context.TODO(), err)
		}
		logger.Info(context.TODO(), "Methods:", strings.Join(methods, ","))
		return nil
		// return err
		/*
			pathRegexp, err := route.GetPathRegexp()
			if err == nil {
				logger.Info("Path regexp:", pathRegexp)
			}
			queriesTemplates, err := route.GetQueriesTemplates()
			if err == nil {
				logger.Info("Queries templates:", strings.Join(queriesTemplates, ","))
			}
			queriesRegexps, err := route.GetQueriesRegexp()
			if err == nil {
				logger.Info("Queries regexps:", strings.Join(queriesRegexps, ","))
			}
		*/
	})
	if err != nil {
		logger.Error(context.TODO(), err)
	}
}

// REGISTERING HOOKS FROM APPLICATION
type InitOption func(*initConfig)

type initConfig struct {
	logging foundation.LoggingOptions
}

func WithLogHook(h logrus.Hook) InitOption {
	return func(cfg *initConfig) {
		cfg.logging.ExtraHooks = append(cfg.logging.ExtraHooks, h)
	}
}
