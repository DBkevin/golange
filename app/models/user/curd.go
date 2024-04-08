package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/route"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
func (user User) Link() string {
	return route.Name2URL("users.show", "id", user.GetStringID())
}
