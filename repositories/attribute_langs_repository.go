package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeLangRepository struct {
	DB *gorm.DB
}

func NewAttributeLangRepository(db *gorm.DB) *AttributeLangRepository {
	return &AttributeLangRepository{
		DB: db,
	}
}
func (r *AttributeLangRepository) CreateAttributeLang(attribute_lang *models.AttributeLang) error {
	return r.DB.Create(attribute_lang).Error
}

func (r *AttributeLangRepository) GetAllAttributeLang() ([]models.AttributeLang, error) {
	var attribute_lang []models.AttributeLang
	if err := r.DB.Find(&attribute_lang).Error; err != nil {
		return nil, err
	}
	return attribute_lang, nil
}
