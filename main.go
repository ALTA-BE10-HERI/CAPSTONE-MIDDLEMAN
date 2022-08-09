package main

import (
	"fmt"
	"middleman-capstone/config"
	"middleman-capstone/factory"
	"middleman-capstone/infrastructure/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.AddTrailingSlash())

	factory.InitFactory(e, db)

	fmt.Println("application is running ...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
