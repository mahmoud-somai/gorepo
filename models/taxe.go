package models

type Tax struct {
	ID        uint `gorm:"primaryKey"`
	TenantID  string
	ForeignID uint
	ShopID    uint
	Rate      string
	Active    uint
	Deleted   uint
}

func (Tax) TableName() string {
	return "taxes"
}

type TaxLang struct {
	ID       uint `gorm:"primaryKey"`
	TenantID string
	TaxID    uint
	LangID   uint
	Name     string
	WooClass string
}

func (TaxLang) TableName() string {
	return "tax_langs"
}
