package easy_serve

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EasyServeConfig struct {
	ConfigFile   string
	CustomConfig interface{}
}

var (
	httpServer *httpserver
	once       = sync.Once{}
)

var configFile = flag.String("i", "config.yaml", "config file")

func New(c *EasyServeConfig) {
	if len(c.ConfigFile) != 0 {
		configFile = &c.ConfigFile
	}
	flag.Parse()
	config.InitConf(*configFile, c.CustomConfig)

	initLogger()
	initRedis()
	initDB()
	initServe()
}

func initServe() {
	for _, ty := range strings.Split(config.C.Service.ServeType, ",") {
		switch ty {
		case config.ServeType_HTTP:
			httpServer = newHttpServer()
		case config.ServeType_CRON:
			cron = *newEasyCron()
		}
	}
}

func Serve() {
	for _, ty := range strings.Split(config.C.Service.ServeType, ",") {
		switch ty {
		case config.ServeType_HTTP:
			httpServer.start()
		case config.ServeType_CRON:
			cron.Start()
		}
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	once.Do(Stop)
}

func Stop() {
	for _, ty := range strings.Split(config.C.Service.ServeType, ",") {
		switch ty {
		case config.ServeType_HTTP:
			httpServer.stop()
		case config.ServeType_CRON:
			cron.Stop()
		}
	}

	for _, client := range redisMap {
		client.close()
	}

	for _, client := range dbMap {
		client.close()
	}
	fmt.Printf("easy serve stop pid: %d\n", os.Getpid())
	os.Exit(0)
}

func HttpServeRouter() *gin.Engine {
	return httpServer.router
}

func CronjobScheduler() *gocron.Scheduler {
	return cron.scheduler
}

func GetRedisClient(name string) (*redis.Client, error) {
	obj, ok := redisMap[name]
	if !ok {
		return nil, fmt.Errorf("name %s not exists", name)
	}
	return obj.client, obj.err
}

func GetGormClient(name string) (*gorm.DB, error) {
	obj, ok := dbMap[name]
	if !ok {
		return nil, fmt.Errorf("name %s not exists", name)
	}
	return obj.db, obj.err
}

func GetLogger(name string) (*zap.Logger, error) {
	logger, ok := loggerMap[name]
	if !ok {
		return nil, fmt.Errorf("name %s not exists", name)
	}
	return logger, nil
}
