package app

import (
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

type Bookmark struct {
	Title         string `json:"title"`
	Link          string `json:"link"`
	SubRedditName string `json:"subRedditName"`
}

func (app *App) saveBookmarks(ctx echo.Context) error {
	var bookmarks []Bookmark
	err := ctx.Bind(&bookmarks)

	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusBadRequest, "Bad request")
	}
	fmt.Println(bookmarks)

	for i, v := range bookmarks {
		_ = i
		fmt.Println(v.Title)
	}

	return ctx.String(http.StatusOK, "Bookmark Saved!")
}
