// address.go

package models

type Addresses struct {
	ID uint `gorm:"primaryKey"`

	Alias        string
	Company      string
	Last_Name    string
	First_Name   string
	VAT_Number   string
	Address1     string
	Address2     string
	Postcode     string
	City         string
	Other        string
	Phone        string
	Phone_Mobile string
	DNI          string
	Woo_Billing  uint
	Woo_Shipping uint
	Woo_Country  string
	Woo_State    string

	Tenant_ID       string
	Shop_ID         uint
	Foreign_ID      uint
	Customer_ID     uint
	Manufacturer_ID uint
	Supplier_ID     uint
	Warehouse_ID    uint
	Country_ID      uint
	State_ID        uint
}

func (Addresses) TableName() string {
	return "addresses"
}
