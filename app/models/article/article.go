package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title      string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body       string `gorm:"type:longtext;not null;" valid:"body"`
	UserID     uint64 `gorm:"not null;index"`
	User       user.User
	CategoryID uint64 `gorm:"not null;index"`
	Category   category.Category
}

func (article Article) Link() string {

	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))

}
