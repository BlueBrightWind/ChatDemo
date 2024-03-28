package middleware

import (
	"ChatDemo/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			common.FailWithDetail(c, "权限不足", nil)
			c.Abort()
			return
		}
		// 提取token的有效部分
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		// 非法token
		if err != nil || !token.Valid {
			common.FailWithDetail(c, "权限不足", nil)
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		UserName := claims.UserName
		// 将用户信息写入上下文便于读取
		c.Set("userId", userId)
		c.Set("userName", UserName)
		c.Next()
	}
}

func SocketAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Query("token")
		// 非法token
		if tokenString == "" {
			common.FailWithDetail(c, "权限不足", nil)
			c.Abort()
			return
		}
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		// 非法token
		if err != nil || !token.Valid {
			common.FailWithDetail(c, "权限不足", nil)
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		UserName := claims.UserName
		// 将用户信息写入上下文便于读取
		c.Set("userId", userId)
		c.Set("userName", UserName)
		c.Next()
	}
}
