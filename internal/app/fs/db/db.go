package db

import (
	"fulfillment-service/internal/app/fs/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

// InitDB initializes the database connection.
func InitDB(dataSourceName string) {
	var err error
	database, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = database.AutoMigrate(&models.DeliveryExecutive{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func GetDB() *gorm.DB {
	return database
}
