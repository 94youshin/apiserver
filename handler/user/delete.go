package user

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler"
	"github.com/youshintop/apiserver/model"
	"github.com/youshintop/apiserver/pkg/errno"
	"strconv"
)

// Delete a user by the user identifier.
func Delete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err = model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
