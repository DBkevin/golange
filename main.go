package main

import (
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"net/http"
)

func init() {
	config.Initialize()
}
func main() {
	// 初始化 SQL

	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(c.GetString("app.host")+":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))

}
