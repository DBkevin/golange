// Package bootstrap 负责应用初始化相关工作，比如初始化路由。
package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/routes"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupDB() {

	db := model.ContentDB()
	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	migration(db)
}
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)

	return router
}

func migration(db *gorm.DB) {

	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
