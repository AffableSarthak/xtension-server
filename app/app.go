package app

import (
	"affableSarthak/extension/server/database"
	"affableSarthak/extension/server/market"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	App struct {
		e  *echo.Echo
		db *gorm.DB
		mp string // KITE | GROWWW
	}
)

func NewApp(echoApp *echo.Echo) App {
	db := database.PgConnect()

	// Based on some value, make it kite or groww
	// Create a helper function to give you all kite or groww related functionality based on the market provider type
	kite := market.NewKiteMarket()
	mp := market.NewMarketProvider(kite)

	fmt.Println(mp)

	app := App{
		e:  echoApp,
		db: db,
		mp: "KITE",
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
