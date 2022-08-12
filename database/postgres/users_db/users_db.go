package users_db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	postgresUsersDbUser     = "POSTGRES_USERS_DB_USER"
	postgresUsersDbPassword = "POSTGRES_USERS_DB_PASSWORD"
	postgresUsersDbHost     = "POSTGRES_USERS_DB_HOST"
	postgresUsersDbName     = "POSTGRES_USERS_DB_NAME"
	postgresUsersDbPort     = "POSTGRES_USERS_DB_PORT"
)

var (
	Client     *gorm.DB
	dbUsername string
	dbPassword string
	dbHost     string
	dbName     string
	dbPort     string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("The .env cannot be loaded")
	}

	dbUsername = os.Getenv(postgresUsersDbUser)
	dbPassword = os.Getenv(postgresUsersDbPassword)
	dbHost = os.Getenv(postgresUsersDbHost)
	dbName = os.Getenv(postgresUsersDbName)
	dbPort = os.Getenv(postgresUsersDbPort)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("The database cannot be open.")
	}
	Client = db
}
