package socket

import (
	"ChatDemo/service"

	"github.com/gin-gonic/gin"
)

func RegisterSocketRouter(r *gin.Engine, middleware gin.HandlerFunc) {
	r.GET("/connect", middleware, service.Connect)
}
