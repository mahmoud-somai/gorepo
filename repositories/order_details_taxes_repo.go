package repositories

import (
	"log"
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type OrderDetailsTaxRepository struct {
	DB *gorm.DB
}

func NewOrderDetailsTaxRepository(db *gorm.DB) *OrderDetailsTaxRepository {
	return &OrderDetailsTaxRepository{DB: db}
}

func (r *OrderDetailsTaxRepository) CreateOrderDetailsTax(orderDetailsTax *models.OrderDetailsTax) error {
	result := r.DB.Create(orderDetailsTax)
	if result.Error != nil {
		log.Printf("Error creating OrderDetailsTax: %v", result.Error)
		return result.Error
	}
	log.Printf("Successfully created OrderDetailsTax: %v", orderDetailsTax)
	return nil
}

func (r *OrderDetailsTaxRepository) GetOrderDetailIDByForeignID(foreignID int) (int, error) {
	var orderDetail models.Orderdetail
	result := r.DB.Where("foreign_id = ?", foreignID).First(&orderDetail)
	if result.Error != nil {
		log.Printf("Error retrieving OrderDetail by foreignID %d: %v", foreignID, result.Error)
		return 0, result.Error
	}
	log.Printf("Successfully retrieved OrderDetail ID %d for foreignID %d", orderDetail.ID, foreignID)
	return int(orderDetail.ID), nil
}

func (r *OrderDetailsTaxRepository) GetTaxIDByForeignID(foreignID int) (int, error) {
	var tax models.Tax
	result := r.DB.Where("foreign_id = ?", foreignID).First(&tax)
	if result.Error != nil {
		log.Printf("Error retrieving Tax by foreignID %d: %v", foreignID, result.Error)
		return 0, result.Error
	}
	log.Printf("Successfully retrieved Tax ID %d for foreignID %d", tax.ID, foreignID)
	return int(tax.ID), nil
}
