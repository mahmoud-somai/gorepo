package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CategoryShopRepository struct {
	DB *gorm.DB
}

func NewCategoryShopRepository(db *gorm.DB) *CategoryShopRepository {
	return &CategoryShopRepository{
		DB: db,
	}
}
func (r *CategoryShopRepository) CreateCategoryShop(category_Shop *models.CategoryShop) error {
	return r.DB.Create(category_Shop).Error
}

func (r *CategoryShopRepository) GetAllCategoryShop() ([]models.CategoryShop, error) {
	var category_Shop []models.CategoryShop
	if err := r.DB.Find(&category_Shop).Error; err != nil {
		return nil, err
	}
	return category_Shop, nil
}
