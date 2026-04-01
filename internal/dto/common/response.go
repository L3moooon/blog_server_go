package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局统一返回结构体
type Response struct {
	Code int         `json:"code"` // 业务码：0成功，非0失败
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// 失败返回
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
}

// 自定义code失败
func FailWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
