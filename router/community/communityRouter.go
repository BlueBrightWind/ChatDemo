package community

import (
	"ChatDemo/service"

	"github.com/gin-gonic/gin"
)

func CommunityCheckRoleRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/", middleware, service.CreateCommunity)
	r.DELETE("/", middleware, service.DeleteCommunity)
	r.PUT("/", middleware, service.UpdateCommunity)
}

func RegisterCommunityRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	CommunityCheckRoleRouter(r, middleware)
}
