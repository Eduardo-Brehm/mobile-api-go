package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB returns a database connection
func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	//dsn is the connection string for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	//Test the connection using ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
