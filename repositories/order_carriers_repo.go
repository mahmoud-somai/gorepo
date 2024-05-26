package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderCarriersRepository struct {
	DB *gorm.DB
}

func NewOrderCarriersRepository(db *gorm.DB) *OrderCarriersRepository {
	return &OrderCarriersRepository{
		DB: db,
	}
}
func (r *OrderCarriersRepository) CreateOrderCarriers(order_carrier *models.OrderCarrier) error {
	return r.DB.Create(order_carrier).Error
}

func (r *OrderCarriersRepository) GetOrderIDByForeignID(OrderID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", OrderID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}
