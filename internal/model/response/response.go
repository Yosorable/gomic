package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

const (
	ERROR   = 7000
	SUCCESS = 0
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, nil, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, nil, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, nil, "操作失败", c)
}

func FailWithError(err error, c *gin.Context) {
	if errCode, succ := err.(ErrorCode); succ {
		Result(errCode.Code, nil, errCode.Error(), c)
		return
	}
	Result(ERROR, nil, err.Error(), c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, nil, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
