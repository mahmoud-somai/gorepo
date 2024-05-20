package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type FeatureRepository struct {
	DB *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) *FeatureRepository {
	return &FeatureRepository{
		DB: db,
	}
}
func (r *FeatureRepository) CreateFeature(feature *models.Feature) error {
	return r.DB.Create(feature).Error
}

func (r *FeatureRepository) CreateFeatureLang(feature_lang *models.FeatureLang) error {
	return r.DB.Create(feature_lang).Error
}

func (r *FeatureRepository) CreateFeatureShop(feature_shop *models.FeatureShop) error {
	return r.DB.Create(feature_shop).Error
}

func (r *FeatureRepository) CreateFeatureValue(feature_value *models.FeatureValue) error {
	return r.DB.Create(feature_value).Error
}

func (r *FeatureRepository) CreateFeatureValueLang(feature_value_lang *models.FeatureValueLang) error {
	return r.DB.Create(feature_value_lang).Error
}
