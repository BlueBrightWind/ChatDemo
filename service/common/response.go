package common

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"` // 0: success, -1: fail
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func FailWithDetail(c *gin.Context, detail string, data any) {
	c.JSON(200, Response{
		Code:    -1,
		Message: detail,
		Data:    data,
	})
}

func SuccessWithDetail(c *gin.Context, detail string, data any) {
	c.JSON(200, Response{
		Code:    0,
		Message: detail,
		Data:    data,
	})
}
