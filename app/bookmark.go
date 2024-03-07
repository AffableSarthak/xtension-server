package app

import (
	"affableSarthak/extension/server/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (app *App) SetupBookmarkGroup() {
	bookmarkGroup := app.e.Group("/bookmarks")
	bookmarkGroup.Use(app.isAuthenticated)

	/// Implement a CRUD for bookmarks - TODO.
	// Create
	bookmarkGroup.POST("/save/:name", app.saveBookmarks)

	// Read
	bookmarkGroup.GET("/:name", app.GetAllBookmarksForUser)
	bookmarkGroup.GET("/preload/:name", app.GetBookmarksUsingPreload)
	// Update

	// Delete

}

type (
	bookmarkRequest struct {
		Title         string `json:"title"`
		Link          string `json:"link"`
		SubRedditName string `json:"subRedditName"`
	}
)

func (app *App) saveBookmarks(ctx echo.Context) error {
	var bookmarks []bookmarkRequest
	username := ctx.Param("name")

	var user models.User

	// get the user data for the given username
	result := app.db.Where("user_name = ?", username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusBadRequest, "Invalid username provided")
	}

	err := ctx.Bind(&bookmarks)

	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	fmt.Println(bookmarks)

	// TODO : lenght conditon
	BookmarkDTO := make([]models.Bookmark, 4, 6)

	for i, bookmark := range bookmarks {
		// DB function to save bookmarks
		fmt.Println("Bookmark", bookmark, i)

		BookmarkDTO[i].Link = bookmark.Link
		BookmarkDTO[i].SubRedditName = bookmark.SubRedditName
		BookmarkDTO[i].Title = bookmark.Title
		BookmarkDTO[i].UserID = user.ID

		res := app.db.Create(&BookmarkDTO[i])
		if res.Error != nil {
			fmt.Println(res.Error)

		}

	}

	var bookmark []models.Bookmark
	// Get bookmarks for user
	res := app.db.Where("user_id = ?", user.ID).Find(&bookmark)
	if res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "Error getting the bookmarks")
	}
	return ctx.JSON(http.StatusOK, map[string][]models.Bookmark{
		"bookmarks": bookmark,
	})
}

func (app *App) GetAllBookmarksForUser(ctx echo.Context) error {
	// Get the username from the params
	username := ctx.Param("name")
	var user models.User
	res := app.db.Where("user_name = ?", username).First(&user)
	if res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "The user doesn't exist")
	}

	var bookmarks []models.Bookmark
	result := app.db.Where("user_id = ?", user.ID).Find(&bookmarks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusOK, map[string]string{
				"bookmarks": "none found",
			})
		}

		return ctx.String(http.StatusInternalServerError, "Error fetching bookmarks")
	}

	return ctx.JSON(http.StatusOK, map[string][]models.Bookmark{
		"bookmarks": bookmarks,
	})
}

func (app *App) GetBookmarksUsingPreload(ctx echo.Context) error {
	username := ctx.Param("name")

	var user models.User

	result := app.db.Where("user_name = ?", username).Preload("Bookmark").Find(&user)

	if result.Error != nil {
		return ctx.String(http.StatusInternalServerError, "User invalid")
	}

	return ctx.JSON(http.StatusOK, map[string][]models.Bookmark{
		"bookmarks": user.Bookmark,
	})
}
