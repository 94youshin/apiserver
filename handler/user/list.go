package user

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler"
	"github.com/youshintop/apiserver/pkg/errno"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.BindJSON(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	//infos, count, err :=
}
