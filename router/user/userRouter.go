package user

import (
	"ChatDemo/service"

	"github.com/gin-gonic/gin"
)

func registerUserNoCheckRoleRouter(r *gin.RouterGroup) {
	r.GET("/isUserNameValid", service.IsUserNameValid)
	r.POST("/", service.CreateUser)
	r.POST("/login", service.Login)
}

func registerUserCheckRoleRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.GET("/", middleware, service.GetUser)
	r.DELETE("/", middleware, service.DeleteUser)
	r.PUT("/", middleware, service.UpdateUserInfo)
}

func RegisterUserRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	registerUserNoCheckRoleRouter(r)
	registerUserCheckRoleRouter(r, middleware)
}
