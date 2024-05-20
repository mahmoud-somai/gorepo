package models

type AttributeShop struct {
	ID           uint `gorm:"primaryKey"`
	Tenant_ID    string
	Attribute_ID uint
	Shop_ID      uint
	Foreign_ID   uint
}

func (AttributeShop) TableName() string {
	return "attribute_shops"
}
