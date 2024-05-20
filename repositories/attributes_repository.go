package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeRepository struct {
	DB *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) *AttributeRepository {
	return &AttributeRepository{
		DB: db,
	}
}
func (r *AttributeRepository) CreateAttribute(attribute *models.Attributes) error {
	return r.DB.Create(attribute).Error
}

func (r *AttributeRepository) GetAllAttribute() ([]models.Attributes, error) {
	var attribute []models.Attributes
	if err := r.DB.Find(&attribute).Error; err != nil {
		return nil, err
	}
	return attribute, nil
}
