package user

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler"
	"github.com/youshintop/apiserver/model"
	"github.com/youshintop/apiserver/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	u, err := model.Get(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, u)
}
