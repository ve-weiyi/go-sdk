package response

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/go-sdk/utils/convert"
	"net/http"
)

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 504
	SUCCESS = 0
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	obj := Response{
		code,
		msg,
		data,
	}

	// 开始时间
	c.String(http.StatusOK, convert.ObjectToJsonSnake(obj))
}

func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMessage(c *gin.Context, message string) {
	Result(c, SUCCESS, message, map[string]interface{}{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "查询成功", data)
}

func OkWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(c, SUCCESS, message, data)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, "操作失败", map[string]interface{}{})
}

func FailWithMessage(c *gin.Context, message string) {
	Result(c, ERROR, message, map[string]interface{}{})
}

func FailWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(c, ERROR, message, data)
}
