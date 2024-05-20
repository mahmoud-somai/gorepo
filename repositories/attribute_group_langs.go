package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type AttributeGroupLangRepository struct {
	DB *gorm.DB
}

func NewAttributeGroupLangRepository(db *gorm.DB) *AttributeGroupLangRepository {
	return &AttributeGroupLangRepository{
		DB: db,
	}
}
func (r *AttributeGroupLangRepository) CreateAttributeGroupLang(attribute_group_lang *models.AttributeGroupLang) error {
	return r.DB.Create(attribute_group_lang).Error
}

func (r *AttributeGroupLangRepository) GetAllAttributeGroupLang() ([]models.AttributeGroupLang, error) {
	var attribute_group_lang []models.AttributeGroupLang
	if err := r.DB.Find(&attribute_group_lang).Error; err != nil {
		return nil, err
	}
	return attribute_group_lang, nil
}
