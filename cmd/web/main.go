package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/snowflake_demo/internal/db"
	"github.com/mariogarzac/snowflake_demo/internal/handlers"
	"github.com/mariogarzac/snowflake_demo/pkg/config"
)

var app config.Config

func main() {

	database := db.NewConnection()

	app.Database = database

	repo := handlers.NewRepo(&app)
	repo.NewRepository(repo)

	e := echo.New()

	// routes(e, h)

	e.Logger.Fatal(e.Start(":3000"))

	defer database.Database.Close()
}
