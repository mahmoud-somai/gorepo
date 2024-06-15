package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AddressesRepository struct {
	DB *gorm.DB
}

func NewAddressesRepository(db *gorm.DB) *AddressesRepository {
	return &AddressesRepository{
		DB: db,
	}
}
func (r *AddressesRepository) CreateAddress(address *models.Addresses) error {
	return r.DB.Create(address).Error
}
