package easy_serve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/gin-gonic/gin"
)

type httpserver struct {
	router *gin.Engine
	server http.Server
}

func newHttpServer() *httpserver {
	router := gin.New()
	router.Use(gin.Recovery())

	if config.C.Service.Env == config.ServeEnv_Pro {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()

	return &httpserver{
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", config.C.Service.ServePort),
			Handler: router,
		},
		router: router,
	}
}

func (s *httpserver) start() error {
	var err error
	go func() {
		err = s.server.ListenAndServe()
	}()
	return err

}

func (s *httpserver) stop() {
	err := s.server.Shutdown(context.Background())
	if err != nil {
		fmt.Printf("httpserver exit failed, err: %v\n", err)
	} else {
		fmt.Println("httpserver exited")
	}
}
