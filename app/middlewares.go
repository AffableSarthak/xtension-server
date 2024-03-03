package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get cookie from request
		cookie, err := c.Request().Cookie("session_id")
		if err != nil {
			if err == http.ErrNoCookie {
				// No cookie found, return unauthorized error
				return echo.ErrUnauthorized
			}
			// Other errors, handle as needed
			return err
		}

		// Check if cookie value is not empty
		if cookie.Value == "" {
			return echo.ErrUnauthorized
		}

		// User is authenticated, call next handler
		return next(c)
	}
}
