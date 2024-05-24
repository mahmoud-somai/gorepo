package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) GetAllOrder() ([]models.Order, error) {
	var order []models.Order
	if err := r.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) CreateOrderDetail(orderDetail *models.Orderdetail) error {
	return r.DB.Create(orderDetail).Error
}
