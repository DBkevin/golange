// Package category 存放应用的分类数据模型
package category

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// caregory文章分类
type Category struct {
	models.BaseModel
	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

// Link方法用来生成文章链接
func (category Category) Link() string {

	return route.Name2URL("home", "id", category.GetStringID())
}
