package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type ShippingRepository struct {
	DB *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *ShippingRepository {
	return &ShippingRepository{
		DB: db,
	}
}

func (r *ShippingRepository) CreateShipping(shipping *models.ShippingAddressesWoo) error {
	return r.DB.Create(shipping).Error
}

func (r *ShippingRepository) GetOrderIDByForeignID(OrderID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", OrderID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}
