package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id       uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreateAt time.Time `gorm:"column:createAt" json:"-"`
	UpdateAt time.Time `gorm:"column:updateAt" json:"-"`
	DeleteAt time.Time `gorm:"column:deleteAt" json:"-"`
}

type UserInfo struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	SayHello string `json:"sayHello"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
