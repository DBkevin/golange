package user

import (
	"goblog/app/models"
)

type User struct {
	models.BaseModel
	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);default:null;unique" valid:"email"`
	Password        string `gorm:"type:varchar(255);not null;" valid:"password"`
	PasswordConfirm string `goro:"-" valid:"password_confirm"`
}
