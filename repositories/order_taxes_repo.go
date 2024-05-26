package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderTaxesRepository struct {
	DB *gorm.DB
}

func NewOrderTaxRepository(db *gorm.DB) *OrderTaxesRepository {
	return &OrderTaxesRepository{
		DB: db,
	}
}
func (r *OrderTaxesRepository) CreateOrderTax(order_tax *models.OrderTax) error {
	return r.DB.Create(order_tax).Error
}

func (r *OrderTaxesRepository) GetOrderIDByForeignID(OrderID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", OrderID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}

func (r *OrderTaxesRepository) GetTaxIDByForeignID(TaxID int) (int, error) {
	var tax models.Tax
	result := r.DB.Where("foreign_id = ?", TaxID).First(&tax)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(tax.ID), nil
}
