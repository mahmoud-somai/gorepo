package migration

import (
	"log"
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

// MigrateDB performs automatic migration for all models
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Billing_addresses_woo{}); err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
		return err
	}
	// Add more models as needed

	log.Println("Database migration completed")
	return nil
}
