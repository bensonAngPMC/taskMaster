package helper

import (
	"net/http"
	"taskMaster/data/response"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func HandleError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, response.Response{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   message,
	})
	ctx.Abort() // 终止后续操作
}
