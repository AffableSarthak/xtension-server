package app

import (
	"affableSarthak/extension/server/models"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (app *App) SetupLoginRoutes() {

	app.e.POST("/login", app.loginHandler)
	app.e.POST("/xtension/login", app.xtensionLoginHandler)
	app.e.POST("/api/v2/login", app.loginV2)
}

func (app *App) loginV2(ctx echo.Context) error {
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")

	hashedPass, err := app.HashPassword(password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Server failed you by not being able to parse your password")
	}

	// DB stuff
	user, err, isNewUser := app.userFirstOrCreate(name, hashedPass)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Error fetching or creating user")
	}

	if isNewUser {
		userSessionId, err := app.cookieLyf(user.SessionID)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		setCookie(ctx, "session_id", userSessionId)

		return ctx.String(http.StatusOK, "Login successful")
	} else {
		// check if entered password is correct
		isCorrectPasskey := app.CheckPasswordHash(password, user.Hp)

		// Password is correct
		if isCorrectPasskey {
			userSessionId, err := app.cookieLyf(user.SessionID)

			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
			setCookie(ctx, "session_id", userSessionId)

			return ctx.String(http.StatusOK, "Login successful")
		} else {
			// Invalid password
			return ctx.String(http.StatusBadRequest, "Either username or password is incorrect")
		}
	}

}

func (app *App) userFirstOrCreate(name, password string) (models.User, error, bool) {
	var user models.User

	result := app.db.Where("user_name = ?", name).Attrs(models.User{
		UserName: name,
		Hp:       password,
		Session: models.Session{
			UserSessionID: uuid.NewString(),
		},
	}).FirstOrCreate(&user)

	if result.Error != nil {
		return user, result.Error, false
	}

	// If no rows were affected then user existed
	if result.RowsAffected == 0 {
		return user, nil, false
	} else {
		return user, nil, true
	}

}

func (app *App) xtensionLoginHandler(ctx echo.Context) error {
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")

	var user models.User
	result := app.db.Where("user_name = ?", name).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Handle record not found
		// This basically means, no user with the name : ${userName}
		// Create user with session
		hashedPass, err := app.HashPassword(password)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Server failed you by not being able to parse your password")
		}

		user := models.User{
			UserName: name,
			Hp:       hashedPass,
			Session: models.Session{
				UserSessionID: uuid.NewString(),
			},
		}

		res := app.db.Create(&user)

		if res.Error != nil {
			// Some DB error in creating user
			return ctx.String(http.StatusInternalServerError, "Sever failed in creating user")
		}

		return ctx.String(http.StatusOK, name)
	}

	if result.Error != nil {
		// Some unknowed DB error occured
		return ctx.String(http.StatusInternalServerError, "DB failure")
	}

	// USER DOES EXIST

	// check if entered password is correct
	isCorrectPasskey := app.CheckPasswordHash(password, user.Hp)

	// Password is correct
	if isCorrectPasskey {
		return ctx.String(http.StatusOK, user.UserName)
	} else {
		// Invalid password
		return ctx.String(http.StatusBadRequest, "Either username or password is incorrect")
	}

}

func (app *App) loginHandler(ctx echo.Context) error {
	userName := ctx.FormValue("userName")
	passKey := ctx.FormValue("passKey")

	var user models.User
	result := app.db.Where("user_name = ?", userName).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Handle record not found
		// This basically means, no user with the name : ${userName}
		// Create user with session
		hashedPass, err := app.HashPassword(passKey)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Server failed you by not being able to parse your password")
		}

		user := models.User{
			UserName: userName,
			Hp:       hashedPass,
			Session: models.Session{
				UserSessionID: uuid.NewString(),
			},
		}

		res := app.db.Create(&user)

		if res.Error != nil {
			// Some DB error in creating user
			return ctx.String(http.StatusInternalServerError, "Sever failed in creating user")
		}

		userSessionId, err := app.cookieLyf(user.SessionID)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		setCookie(ctx, "session_id", userSessionId)

		return ctx.String(http.StatusOK, "Login successful")
	}

	if result.Error != nil {
		// Some unknowed DB error occured
		return ctx.String(http.StatusInternalServerError, "DB failure")
	}

	// USER DOES EXIST

	// check if entered password is correct
	isCorrectPasskey := app.CheckPasswordHash(passKey, user.Hp)

	// Password is correct
	if isCorrectPasskey {
		userSessionId, err := app.cookieLyf(user.SessionID)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		setCookie(ctx, "session_id", userSessionId)

		return ctx.String(http.StatusOK, "Login successful")
	} else {
		// Invalid password
		return ctx.String(http.StatusBadRequest, "Either username or password is incorrect")
	}

}

func (app *App) cookieLyf(sessionID uint) (string, error) {
	var session models.Session
	res := app.db.Where("id = ?", sessionID).First(&session)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err := errors.New("error something unexpected, no session for existing user")
		return "", err
	}
	if res.Error != nil {
		err := errors.New("some error occured getting user session")
		return "", err
	}

	return session.UserSessionID, nil
}

func (app *App) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (app *App) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func setCookie(ctx echo.Context, name, value string) {
	log.Println("Setting Cookie")
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	log.Println(cookie)
	ctx.SetCookie(cookie)
}
