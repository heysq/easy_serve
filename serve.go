package easy_serve

import (
	"flag"
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/Sunqi43797189/easy_serve/serve"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
    "gorm.io/gorm"
)

var (
	httpServer *serve.HttpServer
)

var configFile = flag.String("i", "config.yaml", "config file")

func New() {
	flag.Parse()
	config.InitConf(*configFile)

	initRedis()
	initDB()
	initServe()
}

func initServe() {
	switch config.C.Service.ServeType {
	case config.ServeType_HTTP:
		httpServer = serve.NewHttpServer()
	}
}

func Serve() {
	var err error
	switch config.C.Service.ServeType {
	case config.ServeType_HTTP:
		err = httpServer.Start()
	}
}

func Stop() {
	switch config.C.Service.ServeType {
	case config.ServeType_HTTP:
		httpServer.Stop()
	}
}

func HttpServeRouter() *gin.Engine {
	return httpServer.HttpRouter()
}

func GetRedisClient(name string) (*redis.Client, error){
	obj, ok := redisMap[name]
	if !ok {
		return nil, fmt.Errorf("name %s not exists", name)
	}
	return obj.client, obj.err
}

func GetGormClient(name string) (*gorm.DB, error){
	obj, ok := dbMap[name]
	if !ok {
		return nil, fmt.Errorf("name %s not exists", name)
	}
	return obj.db, obj.err
}