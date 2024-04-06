package user

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (user *User) Create() (err error) {
	fmt.Println("进入create")
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
