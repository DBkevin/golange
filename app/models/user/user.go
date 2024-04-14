package user

import (
	"goblog/app/models"
	"goblog/pkg/model"
	"goblog/pkg/password"
	"goblog/pkg/types"
)

type User struct {
	models.BaseModel
	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);default:null;unique" valid:"email"`
	Password        string `gorm:"type:varchar(255);not null;" valid:"password"`
	PasswordConfirm string `gorm:"-;" valid:"password_confirm"`
}

// Get 通过 ID 获取用户
func Get(idstr string) (User, error) {
	var user User
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_pw string) bool {
	return password.CheckHash(_pw, user.Password)
}
