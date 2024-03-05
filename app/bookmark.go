package app

import (
	"affableSarthak/extension/server/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) SetupBookmarkGroup() {
	bookmarkGroup := app.e.Group("/bookmarks")
	bookmarkGroup.Use(app.isAuthenticated)

	/// Implement a CRUD for bookmarks - TODO.
	// Create
	bookmarkGroup.POST("/save", app.saveBookmarks)

	// Read

	// Update

	// Delete

}

func (app *App) saveBookmarks(ctx echo.Context) error {
	var bookmarks []models.Bookmark
	err := ctx.Bind(&bookmarks)

	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusBadRequest, "Bad request")
	}
	fmt.Println(bookmarks)

	for _, bookmark := range bookmarks {
		// DB function to save bookmarks

		fmt.Println("Bookmark", bookmark)
		result := app.db.Create(&bookmark)

		if result.Error != nil {
			return ctx.String(http.StatusInternalServerError, "Error saving bookmark")
		}

	}

	return ctx.String(http.StatusOK, "Bookmark Saved!")
}
