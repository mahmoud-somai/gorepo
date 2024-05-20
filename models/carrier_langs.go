// carrier_lang.go

package models

type CarrierLang struct {
	ID         uint `gorm:"primaryKey"`
	Carrier_ID uint
	Lang_ID    uint
	Delay      string
}

func (CarrierLang) TableName() string {
	return "carrier_langs"
}
