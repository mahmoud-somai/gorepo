package prestashop_DTOs
import (
    "shifti-connector-backend/models"
)
// ExampleDTO represents a data transfer object for the Example model.
type ExampleDTO struct {
    Namee string
    // Add other fields as needed
}

// ToModel converts the ExampleDTO instance to a Example model instance.
func (dto *ExampleDTO) ToModel() *models.Example {
    return &models.Example{
        Name: dto.Namee,
        // Map other fields as needed
    }
}
