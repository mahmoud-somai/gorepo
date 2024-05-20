package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CarrierShopRepository struct {
	DB *gorm.DB
}

func NewCarrierShopRepository(db *gorm.DB) *CarrierShopRepository {
	return &CarrierShopRepository{
		DB: db,
	}
}
func (r *CarrierShopRepository) CreateCarrierShop(carrier_Shop *models.CarrierShop) error {
	return r.DB.Create(carrier_Shop).Error
}

func (r *CarrierShopRepository) GetAllCarrierShop() ([]models.CarrierShop, error) {
	var carrier_Shop []models.CarrierShop
	if err := r.DB.Find(&carrier_Shop).Error; err != nil {
		return nil, err
	}
	return carrier_Shop, nil
}
