package user

import "github.com/youshintop/apiserver/model"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   string `json:"offset"`
	Limit    string `json:"limit"`
}

type ListResponse struct {
	Total    uint64            `json:"total"`
	UserList []*model.UserInfo `json:"userList"`
}
