package models

type CarrierShop struct {
	ID         uint `gorm:"primaryKey"`
	Carrier_ID uint
	Shop_ID    uint
	Foreign_ID uint
}

func (CarrierShop) TableName() string {
	return "carrier_shops"
}
