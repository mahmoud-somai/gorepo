package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeGroupShopRepository struct {
	DB *gorm.DB
}

func NewAttributeGroupShopRepository(db *gorm.DB) *AttributeGroupShopRepository {
	return &AttributeGroupShopRepository{
		DB: db,
	}
}
func (r *AttributeGroupShopRepository) CreateAttributeGroupShop(attribute_group_shop *models.AttributeGroupShop) error {
	return r.DB.Create(attribute_group_shop).Error
}

func (r *AttributeGroupShopRepository) GetAllAttributeGroupShop() ([]models.AttributeGroupShop, error) {
	var attribute_group_shop []models.AttributeGroupShop
	if err := r.DB.Find(&attribute_group_shop).Error; err != nil {
		return nil, err
	}
	return attribute_group_shop, nil
}
