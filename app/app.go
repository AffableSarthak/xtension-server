package app

import (
	"affableSarthak/extension/server/database"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	App struct {
		e  *echo.Echo
		db *gorm.DB
	}
)

func NewApp(echoApp *echo.Echo) App {
	db := database.PgConnect()

	app := App{
		e:  echoApp,
		db: db,
	}

	app.AppInit()
	return app
}

func (app *App) AppInit() {
	app.SetupLoginRoutes()
	app.SetupUserGroup()
	app.SetupBookmarkGroup()
}

func (app *App) Start() {
	err := app.e.Start(":6969")
	if err == nil {
		log.Fatal("Error starting the server")
	}
}
