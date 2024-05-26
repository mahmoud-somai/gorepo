package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderDetailsRepository struct {
	DB *gorm.DB
}

func NewOrderDetailsRepository(db *gorm.DB) *OrderDetailsRepository {
	return &OrderDetailsRepository{DB: db}
}

func (r *OrderDetailsRepository) CreateOrderDetail(orderDetail *models.Orderdetail) error {
	result := r.DB.Create(orderDetail)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrderDetailsRepository) GetProductIDByForeignID(foreignID int) (int, error) {
	var product models.Product
	result := r.DB.Where("product_id = ?", foreignID).First(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(product.ID), nil
}

func (r *OrderDetailsRepository) GetOrderIDByForeignID(foreignID int) (int, error) {
	var order models.Order
	result := r.DB.Where("foreign_id = ?", foreignID).First(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(order.ID), nil
}
