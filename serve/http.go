package serve

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpServer struct{
	router *gin.Engine
	server http.Server
}


func newHttpServer() *httpServer {
	router := gin.New()
	router.Use(gin.Recovery())

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	return &httpServer{
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", 80),
			Handler: router,
		},
		router: router,
	}
}

func InitHttpServer(engine gin.Engine) {
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}


