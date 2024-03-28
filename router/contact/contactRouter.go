package contact

import (
	"ChatDemo/service"

	"github.com/gin-gonic/gin"
)

func registerContactCheckRoleRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/friend", middleware, service.AddContactFriend)
	r.DELETE("/friend", middleware, service.RemoveContactFriend)
	r.GET("/friend", middleware, service.SearchFriend)

	r.POST("/community", middleware, service.AddContactCommunity)
	r.DELETE("/community", middleware, service.RemoveContactCommunity)
	r.GET("/community", middleware, service.SearchCommunity)

	r.GET("/list", middleware, service.GetContactList)
}

func RegisterContactRouter(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	registerContactCheckRoleRouter(r, middleware)
}
