package db

import (
	"gorm.io/gorm"
)

type Database interface {
	OpenConnection() (*gorm.DB, error)
}
