package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

func Get(idstr string) (Article, error) {

	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.Preload("User").Preload("Category").Find(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetByCategoryID
func GetByCategoryID(cid string) ([]Article, error) {

	var articles []Article
	if err := model.DB.Where("category_id =?", cid).Preload("Category").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// GetByUserID 获取用户全部文章
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int, idr ...[]string) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("home"), perPage)
	// 2. 获取视图数据
	viewData := _pager.Paging()
	// 3. 获取数据

	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil

}
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
func (article *Article) Update() (rowsAffected int64, err error) {
	//result := model.DB.Save(&article)
	//result := model.DB.UpdateColumns(article)
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	resulte := model.DB.Delete(&article)

	if err = resulte.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return resulte.RowsAffected, nil
}
