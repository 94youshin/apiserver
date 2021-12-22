package user

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler"
	"github.com/youshintop/apiserver/model"
	"github.com/youshintop/apiserver/pkg/errno"
	"github.com/youshintop/apiserver/pkg/util"
	"github.com/youshintop/log"
	"strconv"
)

func Update(c *gin.Context) {
	log.Infow("Update function called.", "X-Request-Id", util.GetRequestID(c))

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	var u model.UserModel
	if err = c.BindJSON(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u.Id = uint64(userId)

	// Validate
	if err = u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err = model.Update(&u); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
