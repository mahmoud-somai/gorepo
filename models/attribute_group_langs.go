package models

type AttributeGroupLang struct {
	ID               uint `gorm:"primaryKey"`
	Tenant_ID        string
	AttributeGroupID uint
	Lang_ID          uint
	Name             string
	Public_Name      string
}

func (AttributeGroupLang) TableName() string {
	return "attribute_group_langs"
}
