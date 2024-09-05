package db

import (
	"github.com/SamuelBarbosa1/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
