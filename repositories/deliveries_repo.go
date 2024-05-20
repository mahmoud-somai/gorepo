package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	DB *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{
		DB: db,
	}
}
func (r *DeliveryRepository) CreateDelivery(delivery *models.Delivery) error {
	return r.DB.Create(delivery).Error
}
