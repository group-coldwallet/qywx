package httpresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HttpRespOkOnly(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: GetMsg(SUCCESS),
		Data:    nil,
	})
}

func HttpRespOK(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: GetMsg(SUCCESS),
		Data:    data,
	})
}

func HttpRespErrorOnly(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    FAIL,
		Message: GetMsg(SUCCESS),
		Data:    nil,
	})
}

//httpCode http状态码
//errCode  业务执行码
//msg  返回消息
//data 返回数据
func HttpRespError(ctx *gin.Context, code HttpCode, errMsg string, data interface{}) {
	if code == SUCCESS {
		code = FAIL
	}
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: GetMsg(SUCCESS),
		Data:    data,
	})
}
