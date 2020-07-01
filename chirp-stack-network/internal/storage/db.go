package storage

import "github.com/jmoiron/sqlx"

// db holds the PostgreSQL connection pool.
var db *DBLogger

// DBLogger is a DB wrapper which logs the executed sql queries and their
// duration.
type DBLogger struct {
	*sqlx.DB
}

// DB returns the PostgreSQL database object.
func DB() *DBLogger {
	return db
}


