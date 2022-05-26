package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	DB_USER     = "postgres"
// 	DB_PASSWORD = "kreator98"
// 	DB_NAME     = "seniorproject"
// )

// DB set up
func SetupDB() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")

	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbURL)
	CheckErr(err)

	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
