package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresSQLDatabase(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection to ensure it's valid
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
