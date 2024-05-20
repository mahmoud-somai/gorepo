package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		DB: db,
	}
}
func (r *CountryRepository) CreateCountry(country *models.Country) error {
	return r.DB.Create(country).Error
}

func (r *CountryRepository) CreateCountryLang(country_lang *models.CountryLang) error {
	return r.DB.Create(country_lang).Error
}

func (r *CountryRepository) CreateCountryShop(country_shop *models.CountryShop) error {
	return r.DB.Create(country_shop).Error
}

// func (r *CategoryRepository) GetAllCategory() ([]models.Category, error) {
// 	var category []models.Category
// 	if err := r.DB.Find(&category).Error; err != nil {
// 		return nil, err
// 	}
// 	return category, nil
// }
