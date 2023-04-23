package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn, err := initEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to open DB")
	}

}

func initEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	hostEnv := os.Getenv("HOST")
	dbUserEnv := os.Getenv("DBUSER")
	passwordEnv := os.Getenv("PASSWORD")
	portEnv := os.Getenv("PORT")
	sslModEnv := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v sslmode=%v", hostEnv, dbUserEnv, passwordEnv, portEnv, sslModEnv)

	return dsn, nil
}
