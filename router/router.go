package router

import (
	"ChatDemo/middleware"
	"ChatDemo/router/community"
	"ChatDemo/router/contact"
	"ChatDemo/router/socket"
	"ChatDemo/router/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/user")
	user.RegisterUserRouter(userRouter, middleware.AuthMiddleware())

	communityRouter := r.Group("/community")
	community.RegisterCommunityRouter(communityRouter, middleware.AuthMiddleware())

	contactRouter := r.Group("/contact")
	contact.RegisterContactRouter(contactRouter, middleware.AuthMiddleware())

	socket.RegisterSocketRouter(r, middleware.SocketAuthMiddleware())

	return r
}
