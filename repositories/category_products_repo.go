package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CategoryProductRepository struct {
	DB *gorm.DB
}

func NewCategoryProductRepository(db *gorm.DB) *CategoryProductRepository {
	return &CategoryProductRepository{
		DB: db,
	}
}
func (r *CategoryProductRepository) CreateCategoryProduct(category_Product *models.CategoryProduct) error {
	return r.DB.Create(category_Product).Error
}

func (r *CategoryProductRepository) GetAllCategoryProduct() ([]models.CategoryProduct, error) {
	var category_Product []models.CategoryProduct
	if err := r.DB.Find(&category_Product).Error; err != nil {
		return nil, err
	}
	return category_Product, nil
}
