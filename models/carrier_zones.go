package models

type CarrierZone struct {
	ID         uint `gorm:"primaryKey"`
	Carrier_ID uint
	Zone_ID    uint
}

func (CarrierZone) TableName() string {
	return "carrier_zones"
}
