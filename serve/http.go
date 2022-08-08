package serve

import (
	"fmt"
	"net/http"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	router *gin.Engine
	server http.Server
}

func NewHttpServer() *HttpServer {
	router := gin.New()
	router.Use(gin.Recovery())
	
	if config.C.Service.Env != config.ServeEnv_Pro {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()

	return &HttpServer{
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", config.C.Service.ServePort),
			Handler: router,
		},
		router: router,
	}
}

func (s *HttpServer) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *HttpServer) Stop() {

}

func (s *HttpServer) HttpRouter() *gin.Engine {
	return s.router
}
