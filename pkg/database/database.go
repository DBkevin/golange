package database

import (
	"database/sql"
	"goblog/pkg/logger"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	config := mysql.Config{
		User:                    "goblog",
		Passwd:                  "asdasd123123",
		Addr:                    "120.25.70.117:3306",
		Net:                     "tcp",
		DBName:                  "goblog",
		AllowCleartextPasswords: true,
		AllowNativePasswords:    true,
	}
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)
}
func CreateTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
	);`
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)

}
