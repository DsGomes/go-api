package db

import (
	"fmt"
	"log"

	"github.com/dsgomes/rest-api/configs"
	"github.com/dsgomes/rest-api/internal/core/domain"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type database struct{}

func NewPostgres() Database {
	return &database{}
}

func (d *database) OpenConnection() (*gorm.DB, error) {
	conf := configs.GetDB()

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port)

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to the database")
	}
	DB.AutoMigrate(&domain.Todo{})

	return DB, err
}
