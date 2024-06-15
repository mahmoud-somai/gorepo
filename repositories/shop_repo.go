package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type ShopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{DB: db}
}

func (r *ShopRepository) GetShopByURL(url string) (*models.Shop, error) {
	var shop models.Shop
	if err := r.DB.Where("url = ?", url).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}
