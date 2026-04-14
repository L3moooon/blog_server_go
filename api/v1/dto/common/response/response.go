package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageInfo struct {
	Page     int   `json:"page" form:"page"`         // 页码
	PageSize int   `json:"pageSize" form:"pageSize"` // 每页数量
	Total    int64 `json:"total"`                    // 总数
}

// 全局统一返回结构体
type Response struct {
	Code int         `json:"code"` // 业务码：0成功，非0失败
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMessage(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, map[string]interface{}{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "操作成功", data)
}

func OkWithDetailed(c *gin.Context, data interface{}, msg string) {
	Result(c, SUCCESS, msg, data)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, "操作失败", map[string]interface{}{})
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(c, ERROR, msg, map[string]interface{}{})
}

func FailWithDetailed(c *gin.Context, data interface{}, msg string) {
	Result(c, ERROR, msg, data)
}

func NotAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		msg,
		nil,
	})
}
