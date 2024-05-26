package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderFeesRepository struct {
	DB *gorm.DB
}

func NewOrderFeeRepository(db *gorm.DB) *OrderFeesRepository {
	return &OrderFeesRepository{
		DB: db,
	}
}
func (r *OrderFeesRepository) CreateOrderFee(order_fee *models.OrderFee) error {
	return r.DB.Create(order_fee).Error
}

func (r *OrderFeesRepository) GetOrderIDByForeignID(OrderID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", OrderID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}
