package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CarrierRepository struct {
	DB *gorm.DB
}

func NewCarrierRepository(db *gorm.DB) *CarrierRepository {
	return &CarrierRepository{
		DB: db,
	}
}
func (r *CarrierRepository) CreateCarrier(carrier *models.Carrier) error {
	return r.DB.Create(carrier).Error
}

func (r *CarrierRepository) GetAllCarrier() ([]models.Carrier, error) {
	var carrier []models.Carrier
	if err := r.DB.Find(&carrier).Error; err != nil {
		return nil, err
	}
	return carrier, nil
}
