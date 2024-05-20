package models

// Example represents an example entity.
type Example struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
    // Add other fields as needed
}
