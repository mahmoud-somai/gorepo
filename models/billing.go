package models

type Billing_addresses_woo struct {
	ID uint `gorm:"primaryKey"`

	First_Name string
	Last_Name  string
	Company    string
	Address_1  string
	Address_2  string
	City       string
	State      string
	Postcode   string
	Country    string
	Email      string
	Phone      string

	Order_ID uint
}

func (Billing_addresses_woo) TableName() string {
	return "billing_addresses_woo"
}
