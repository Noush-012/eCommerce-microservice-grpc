package db

import (
	"errors"
	"log"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnToDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := cfg.DATABASE

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.SkipDefaultTransaction = true
		log.Println("failed to connect database !")
		return nil, errors.New("failed to connect database ")
	}
	log.Println("Successfully connected to database !")

	db.AutoMigrate()

	return db, err

}
