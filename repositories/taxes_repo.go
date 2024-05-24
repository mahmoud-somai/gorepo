package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type TaxeRepository struct {
	DB *gorm.DB
}

func NewTaxeRepository(db *gorm.DB) *TaxeRepository {
	return &TaxeRepository{
		DB: db,
	}
}

func (r *TaxeRepository) CreateTaxe(taxe *models.Tax) error {
	return r.DB.Create(taxe).Error
}

func (r *TaxeRepository) CreateTaxeLang(taxeLang *models.TaxLang) error {
	return r.DB.Create(taxeLang).Error
}
