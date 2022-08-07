package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func InitHttpServer(engine gin.Engine) {
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}


