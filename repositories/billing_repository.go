package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type BillingRepository struct {
	DB *gorm.DB
}

func NewBillingRepository(db *gorm.DB) *BillingRepository {
	return &BillingRepository{
		DB: db,
	}
}

func (r *BillingRepository) CreateBilling(billing *models.BillingAddressesWoo) error {
	return r.DB.Create(billing).Error
}

func (r *BillingRepository) GetOrderIDByForeignID(OrderID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", OrderID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}
