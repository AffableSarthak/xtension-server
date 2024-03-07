package app

import (
	"affableSarthak/extension/server/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (app *App) SetupUserGroup() {
	userGroup := app.e.Group("/users")
	userGroup.Use(app.isAuthenticated)

	userGroup.GET("/info/:name", app.GetUserInfoByUserName)
}

func (app *App) GetUserInfoByUserName(ctx echo.Context) error {
	username := ctx.Param("name")

	var user models.User
	result := app.db.Where("user_name = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusOK, map[string]struct{}{
				"user": {},
			})
		}

		return ctx.String(http.StatusInternalServerError, "Invalid user or server error error fetching")
	}

	return ctx.JSON(http.StatusOK, map[string]*models.User{
		"data": &user,
	})
}
