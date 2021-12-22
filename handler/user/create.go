package user

import (
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler"
	"github.com/youshintop/apiserver/model"
	"github.com/youshintop/apiserver/pkg/errno"
	"github.com/youshintop/apiserver/pkg/util"
	"github.com/youshintop/log"
)

func Create(c *gin.Context) {

	log.Infow("User create function called.", "X-Request-Id", util.GetRequestID(c))
	var r CreateRequest
	if err := c.BindJSON(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//Validate the data
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := model.Create(&u); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, err)
		return
	}

	handler.SendResponse(c, nil, CreateResponse{Username: r.Username})
}
