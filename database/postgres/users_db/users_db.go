package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_db_user     = "MYSQL_USERS_DB_USER"
	mysql_users_db_password = "MYSQL_USERS_DB_PASSWORD"
	mysql_users_db_host     = "MYSQL_USERS_DB_HOST"
	mysql_users_db_name     = "MYSQL_USERS_DB_NAME"
)

var (
	Client     *sql.DB
	dbUsername = os.Getenv(mysql_users_db_user)
	dbPassword = os.Getenv(mysql_users_db_password)
	dbHost     = os.Getenv(mysql_users_db_host)
	dbName     = os.Getenv(mysql_users_db_name)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		dbUsername,
		dbPassword,
		dbHost,
		dbName,
	)

	var err error
	Client, err = sql.Open("postgresql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connection done.")
}
