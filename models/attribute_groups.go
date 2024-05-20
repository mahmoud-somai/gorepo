package models

type AttributeGroup struct {
	ID             uint `gorm:"primaryKey"`
	Tenant_ID      string
	Type           string
	Position       int
	Order_By       string
	Is_Color_Group uint
}

func (AttributeGroup) TableName() string {
	return "attribute_groups"
}
