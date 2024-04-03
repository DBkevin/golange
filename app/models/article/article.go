package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	ID    uint64
	Title string
	Body  string
}

func (article Article) Link() string {

	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))

}
