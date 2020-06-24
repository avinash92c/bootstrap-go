Usage Guide & Features Guide

- [Configuration Docs](https://github.com/avinash92c/bootstrap-go/blob/master/docs/config/config.md)
- [Diagnostics](https://github.com/avinash92c/bootstrap-go/blob/master/docs/diagnostics/readme.md)

- Subject to update as library develops toward v1 release

```go
import (
	rest "gokit/testappsvcliv/internal/rest"
	cmd "github.com/avinash92c/bootstrap-go"
	bsrest "github.com/avinash92c/bootstrap-go/rest"
)

//Init function
func Init() {
	appsvr, router := cmd.Init()
	routeEngine := bsrest.NewRoutingEngine(router, appsvr)
	rest.DeclareRoutes(routeEngine)
	cmd.StartServer(appsvr, router)
}



package rest

import (
	"fmt"
	"net/http"

	bootstrapmodel "github.com/avinash92c/bootstrap-go/model"
	bootstraprest "github.com/avinash92c/bootstrap-go/rest"
)

func DeclareRoutes(routeEngine bootstraprest.RoutingEngine) {
	routeEngine.BuildRoute("/test", "GET", false, sampleHandler)
}

func sampleHandler(appsvr *bootstrapmodel.AppServer, w http.ResponseWriter, r *http.Request) {
	//DO SOMETHING
	fmt.Fprintf(w, "%v", "Test")
	w.WriteHeader(http.StatusOK)
}

```
