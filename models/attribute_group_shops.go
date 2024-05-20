package models

type AttributeGroupShop struct {
	ID               uint `gorm:"primaryKey"`
	Tenant_ID        string
	AttributeGroupID uint
	Shop_ID          uint
	Foreign_ID       uint
}

func (AttributeGroupShop) TableName() string {
	return "attribute_group_shops"
}
