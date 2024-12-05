package database

import (
	"log"
	"travelnest/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=example password=example dbname=travelnest port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		return nil, err
	}
	
	if err := db.AutoMigrate(&model.Room{}); err != nil {
		log.Fatalf("Error to run migration: %v", err)
	}
	
	return db, nil
}