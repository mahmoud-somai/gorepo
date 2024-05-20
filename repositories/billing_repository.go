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
func (r *BillingRepository) CreateBilling(billing *models.Billing_addresses_woo) error {
	return r.DB.Create(billing).Error
}

func (r *BillingRepository) GetAllBilling() ([]models.Billing_addresses_woo, error) {
	var billing []models.Billing_addresses_woo
	if err := r.DB.Find(&billing).Error; err != nil {
		return nil, err
	}
	return billing, nil
}
