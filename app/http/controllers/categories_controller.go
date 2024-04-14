package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// CategoriesController文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{
		"Category": view.D{
			"Name": "",
		},
	}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	//初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	//2.验证表单
	errors := requests.ValidateCategoryForm(_category)
	// 3.错误检测
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}

}

// Show 显示分类

func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	cid := route.GetRouteVariable("id", r)
	articles, err := article.GetByCategoryID(cid)
	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{"Articles": articles}, "articles.index", "articles._article_meta")
	}
}
