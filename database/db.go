package database

import (
	"database/sql"
	"fmt"
	"log"
)

var Mysql *sql.DB

func Connect() *sql.DB {
	user := "root"
	host := "localhost"
	port := "3306"
	database := "go_auth"

	connection := fmt.Sprintf("%s:@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, host, port, database)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Print(err, "\nError connect database")
	}
	if db == nil {
		panic("db nil")
	}
	migrate(db)
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
		updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp()
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		log.Print(err)
		return
	}
}