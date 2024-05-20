package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CarrierLangRepository struct {
	DB *gorm.DB
}

func NewCarrierLangRepository(db *gorm.DB) *CarrierLangRepository {
	return &CarrierLangRepository{
		DB: db,
	}
}
func (r *CarrierLangRepository) CreateCarrierLang(carrier_Lang *models.CarrierLang) error {
	return r.DB.Create(carrier_Lang).Error
}

func (r *CarrierLangRepository) GetAllCarrierLang() ([]models.CarrierLang, error) {
	var carrier_Lang []models.CarrierLang
	if err := r.DB.Find(&carrier_Lang).Error; err != nil {
		return nil, err
	}
	return carrier_Lang, nil
}
