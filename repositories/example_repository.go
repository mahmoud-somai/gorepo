package repositories

import (
    "gorm.io/gorm"
    "shifti-connector-backend/models"
)

// ExampleRepository represents the repository for Example model.
type ExampleRepository struct {
    DB *gorm.DB
}

// NewExampleController creates a new instance of ExampleController.
func NewExampleRepository(db *gorm.DB) *ExampleRepository {
    return &ExampleRepository{
        DB: db,
    }
}

// CreateExample creates a new example entity in the database.
func (r *ExampleRepository) CreateExample(example *models.Example) error {
    return r.DB.Create(example).Error
}

// GetExampleByID retrieves an example entity from the database by ID.
func (r *ExampleRepository) GetExampleByID(id uint) (*models.Example, error) {
    var example models.Example
    if err := r.DB.First(&example, id).Error; err != nil {
        return nil, err
    }
    return &example, nil
}

// UpdateExample updates an example entity in the database.
func (r *ExampleRepository) UpdateExample(example *models.Example) error {
    return r.DB.Save(example).Error
}

// DeleteExampleByID deletes an example entity from the database by ID.
func (r *ExampleRepository) DeleteExampleByID(id uint) error {
    return r.DB.Delete(&models.Example{}, id).Error
}
