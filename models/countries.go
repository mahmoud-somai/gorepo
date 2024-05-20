package models

type Country struct {
	ID                         uint `gorm:"primaryKey"`
	Tenant_ID                  string
	Zone_ID                    uint
	Currency_ID                uint
	Foreign_ID                 uint
	Call_Prefix                string
	Iso_Code                   string
	Active                     uint
	Contains_States            uint
	Need_Identification_Number uint
	Need_Zip_Code              uint
	Zip_Code_Format            string
	Display_Tax_Label          uint
}

func (Country) TableName() string {
	return "countries"
}

type CountryLang struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Country_ID uint
	Lang_ID    uint
	Name       string
}

func (CountryLang) TableName() string {
	return "country_langs"
}

type CountryShop struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Country_ID uint
	Shop_ID    uint
	Foreign_ID uint
}

func (CountryShop) TableName() string {
	return "country_shops"
}
