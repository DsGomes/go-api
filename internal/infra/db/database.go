package db

import "database/sql"

type Database interface {
	OpenConnection() (*sql.DB, error)
}
