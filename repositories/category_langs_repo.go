package repositories

// import (
// 	"shifti-connector-backend/models"

// 	"gorm.io/gorm"
// )

// type CategoryLangRepository struct {
// 	DB *gorm.DB
// }

// func NewCategoryLangRepository(db *gorm.DB) *CategoryLangRepository {
// 	return &CategoryLangRepository{
// 		DB: db,
// 	}
// }
// func (r *CategoryLangRepository) CreateCategoryLang(category_Lang *models.CategoryLang) error {
// 	return r.DB.Create(category_Lang).Error
// }

// func (r *CategoryLangRepository) GetAllCategoryLang() ([]models.CategoryLang, error) {
// 	var category_Lang []models.CategoryLang
// 	if err := r.DB.Find(&category_Lang).Error; err != nil {
// 		return nil, err
// 	}
// 	return category_Lang, nil
// }
