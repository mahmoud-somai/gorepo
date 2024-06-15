package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) CreateCategoryLang(categoryLang *models.CategoryLang) error {
	return r.DB.Create(categoryLang).Error
}

func (r *CategoryRepository) CreateCategoryShop(categoryShop *models.CategoryShop) error {
	return r.DB.Create(categoryShop).Error
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Check if a category exists by ForeignID and TenantID
func (r *CategoryRepository) CategoryExists(tenantID string, foreignID uint) (bool, error) {
	var category models.CategoryShop
	if err := r.DB.Where("tenant_id = ? AND foreign_id = ?", tenantID, foreignID).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateFullCategory handles the creation of a category, categoryLang, and categoryShop within a transaction
func (r *CategoryRepository) CreateFullCategory(category *models.Category, categoryLang *models.CategoryLang, categoryShop *models.CategoryShop) error {
	// Begin a transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert into categories table
	if err := tx.Create(category).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the CategoryID for categoryLang and categoryShop
	categoryLang.Category_ID = category.ID
	categoryShop.Category_ID = category.ID

	// Insert into categorylangs table
	if err := tx.Create(categoryLang).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert into categoryshops table
	if err := tx.Create(categoryShop).Error; err != nil {
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
