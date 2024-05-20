package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeShopRepository struct {
	DB *gorm.DB
}

func NewAttributeShopRepository(db *gorm.DB) *AttributeShopRepository {
	return &AttributeShopRepository{
		DB: db,
	}
}
func (r *AttributeShopRepository) CreateAttributeShop(attribute_shop *models.AttributeShop) error {
	return r.DB.Create(attribute_shop).Error
}

func (r *AttributeShopRepository) GetAllAttributeShop() ([]models.AttributeShop, error) {
	var attribute_shop []models.AttributeShop
	if err := r.DB.Find(&attribute_shop).Error; err != nil {
		return nil, err
	}
	return attribute_shop, nil
}
