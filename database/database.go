package database

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "seniorproject"
)

// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	CheckErr(err)

	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
