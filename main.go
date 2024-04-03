package main

import (
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"

	"net/http"
)

func main() {
	database.InitDB()
	database.CreateTables()

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe("192.168.0.197:3020", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
