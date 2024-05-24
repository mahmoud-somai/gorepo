package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	return pr.DB.Create(product).Error
}

func (pr *ProductRepository) CreateProductLang(productLang *models.ProductLang) error {
	return pr.DB.Create(productLang).Error
}

func (pr *ProductRepository) CreateProductShop(productShop *models.ProductShop) error {
	return pr.DB.Create(productShop).Error
}

func (pr *ProductRepository) CreateFullProduct(product *models.Product, productLang *models.ProductLang, productShop *models.ProductShop) error {
	tx := pr.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert into products table
	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the ProductID for productLang and productShop
	productLang.ProductID = product.ID
	productShop.ProductID = product.ID

	// Insert into product_langs table
	if err := tx.Create(productLang).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert into product_shops table
	if err := tx.Create(productShop).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
