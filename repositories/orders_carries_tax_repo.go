package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderCarrierTaxRepository struct {
	DB *gorm.DB
}

func NewOrderCarrierTaxRepository(db *gorm.DB) *OrderCarrierTaxRepository {
	return &OrderCarrierTaxRepository{DB: db}
}

func (r *OrderCarrierTaxRepository) CreateOrderCarrierTax(orderCarrierTax *models.OrderCarrierTax) error {
	result := r.DB.Create(orderCarrierTax)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrderCarrierTaxRepository) GetOrderCarrierIDByForeignID(foreignID int) (int, error) {
	var orderCarrier models.OrderCarrier
	result := r.DB.Where("foreign_id = ?", foreignID).First(&orderCarrier)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(orderCarrier.ID), nil
}

func (r *OrderCarrierTaxRepository) GetTaxIDByForeignID(foreignID int) (int, error) {
	var tax models.Tax
	result := r.DB.Where("foreign_id = ?", foreignID).First(&tax)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(tax.ID), nil
}
