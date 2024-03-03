package database

import (
	"affableSarthak/extension/server/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PgConnect() *gorm.DB {
	PG_HOST := os.Getenv("PG_HOST")
	PG_USER := os.Getenv("PG_USER")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_PORT := os.Getenv("PG_PORT")

	//To get here, download the postgres cli,
	// by default it starts a PostgresServer and creates a user postgres and a Database postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=Asia/Kolkata", PG_HOST, PG_USER, PG_PASSWORD, PG_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Apply migration
	err = Migrate(db)
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func Migrate(db *gorm.DB) error {
	// AutoMigrate will create the table if it doesn't exist
	err := db.AutoMigrate(&models.User{}, &models.Session{}, &models.Bookmark{})
	if err != nil {
		return err
	}

	return nil
}
