package easy_serve

import (
	"flag"
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/Sunqi43797189/easy_serve/serve"
	"github.com/gin-gonic/gin"
)

const (
	ServeType_HTTP = "http"
	ServeType_GRPC = "grpc"
)

var (
	httpServer *serve.HttpServer
)

var configFile = flag.String("i", "config.yaml", "config file")

func New() {
	flag.Parse()
	config.InitConf(*configFile)
	fmt.Println(config.C)
	initServe()
}

func initServe() {
	switch config.C.Service.ServeType {
	case ServeType_HTTP:
		httpServer = serve.NewHttpServer()
	}
}

func Serve() {
	var err error
	switch config.C.Service.ServeType {
	case ServeType_HTTP:
		err = httpServer.Start()
	}
	fmt.Println(err.Error())
}

func Stop() {
	switch config.C.Service.ServeType {
	case ServeType_HTTP:
		httpServer.Stop()
	}
}

func HttpServeRouter() *gin.Engine {
	return httpServer.HttpRouter()
}
