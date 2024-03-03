package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) SetupUserGroup() {
	userGroup := app.e.Group("/users")
	userGroup.Use(app.isAuthenticated)

	userGroup.GET("/info/:name", app.GetUserInfoByUserName)
}

func (app *App) GetUserInfoByUserName(ctx echo.Context) error {
	name := ctx.Param("name")
	_ = name
	return ctx.String(http.StatusOK, "lulla")
}
