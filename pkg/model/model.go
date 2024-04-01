package model

import (
	"goblog/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ContentDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "goblog:asdasd123123@tcp(120.25.70.117:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})
	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	logger.LogError(err)
	return DB

}
