package models

type Attributes struct {
	ID uint `gorm:"primaryKey"`

	Foreign_ID uint
	Color      string
	Position   int

	Tenant_ID          string
	Attribute_Group_ID uint
}

func (Attributes) TableName() string {
	return "attributes"
}
