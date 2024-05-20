// attribute_lang.go

package models

type AttributeLang struct {
	ID           uint `gorm:"primaryKey"`
	Tenant_ID    string
	Attribute_ID uint
	Lang_ID      uint
	Name         string
	Description  string
}

func (AttributeLang) TableName() string {
	return "attribute_langs"
}
