package database

import (
	"RestAPI/pkg/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importing the postgres driver
)

var db *sql.DB

func InitDb(config *config.Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DbHost, config.DBPort, config.DbUser, config.DbPassword, config.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func GetDb() *sql.DB {
	return db
}
