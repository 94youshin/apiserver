package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/youshintop/apiserver/pkg/auth"

	"github.com/youshintop/apiserver/pkg/db"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func Create(u *UserModel) error {
	return db.Database().Create(&u).Error
}

func DeleteUser(id uint64) error {
	u := &UserModel{}
	u.BaseModel.Id = id
	return db.Database().Delete(&u).Error
}

func Update(u *UserModel) error {
	return db.Database().Save(&u).Error
}

func Get(username string) (*UserModel, error) {
	u := &UserModel{}
	d := db.Database().Where("username = ?", username).First(&u)
	return u, d.Error
}

func List(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = 50
	}

	var count int64
	users := make([]*UserModel, 0)
	where := fmt.Sprintf("%%%s%%", username)
	if err := db.Database().Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := db.Database().Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}

func (u *UserModel) Compare(password string) error {
	return auth.Compare(u.Password, password)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Validate() error {
	return validator.New().Struct(u)
}
