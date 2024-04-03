package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

// 渲染视图
func Render(w io.Writer, name string, data interface{}) {
	// 1 设置模板相对路径
	viewDir := "resources/views/"
	// 2. 语法糖，将 articles.show 更正为 articles/show
	name = strings.Replace(name, ".", "/", -1)
	// 3 所有布局模板文件 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	//4 在slice里新增我们的目标文件
	newfiles := append(files, viewDir+name+".gohtml")
	//5 解析所有的模板文件
	teml, err := template.New(name + ".gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newfiles...)
	logger.LogError(err)
	//渲染模板
	err = teml.ExecuteTemplate(w, "myapp", data)
	logger.LogError(err)

}
