package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeGroupRepository struct {
	DB *gorm.DB
}

func NewAttributeGroupRepository(db *gorm.DB) *AttributeGroupRepository {
	return &AttributeGroupRepository{
		DB: db,
	}
}
func (r *AttributeGroupRepository) CreateAttributeGroup(attribute_group *models.AttributeGroup) error {
	return r.DB.Create(attribute_group).Error
}

func (r *AttributeGroupRepository) GetAllAttributeGroup() ([]models.AttributeGroup, error) {
	var attribute_group []models.AttributeGroup
	if err := r.DB.Find(&attribute_group).Error; err != nil {
		return nil, err
	}
	return attribute_group, nil
}
