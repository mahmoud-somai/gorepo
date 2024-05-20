package models

type Delivery struct {
	ID              uint `gorm:"primaryKey"`
	Carrier_ID      uint
	Range_Price_ID  uint
	Range_Weight_ID uint
	Zone_ID         uint
	Shop_ID         uint
	Price           float64
	Foreign_ID      uint
}

func (Delivery) TableName() string {
	return "deliveries"
}
