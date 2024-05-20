package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CarrierZoneRepository struct {
	DB *gorm.DB
}

func NewCarrierZoneRepository(db *gorm.DB) *CarrierZoneRepository {
	return &CarrierZoneRepository{
		DB: db,
	}
}
func (r *CarrierZoneRepository) CreateCarrierZone(carrier_Zone *models.CarrierZone) error {
	return r.DB.Create(carrier_Zone).Error
}

func (r *CarrierZoneRepository) GetAllCarrierZone() ([]models.CarrierZone, error) {
	var carrier_Zone []models.CarrierZone
	if err := r.DB.Find(&carrier_Zone).Error; err != nil {
		return nil, err
	}
	return carrier_Zone, nil
}
