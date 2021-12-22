package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/pkg/errno"
	"net/http"
)

type ResponseCode struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSONP(http.StatusOK, ResponseCode{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
