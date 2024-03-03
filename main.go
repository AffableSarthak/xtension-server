package main

import (
	"affableSarthak/extension/server/app"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env := os.Getenv("XTENSION_ENV")

	if env == "" {
		env = "development"
	}

	if env == "production" {
		err := godotenv.Load()

		if err != nil {
			fmt.Println(env)
			log.Fatal("Error loading the Production env")
		}
	} else {

		err := godotenv.Load(".env." + env)

		if err != nil {
			fmt.Println(env)
			// log.Fatal("Error loading the env")
			log.Fatalf("Error loadding the %s env", env)
		}
	}

	e := echo.New()

	// CORS middleware configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true, // Allow cookies to be sent with requests
	}))

	app := app.NewApp(e)

	app.Start()

}
