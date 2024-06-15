// address.go

package models

type Addresses struct {
	ID uint `gorm:"primaryKey"`

	Company         string
	Last_Name       string
	First_Name      string
	Address1        string
	Address2        string
	Postcode        string
	City            string
	Phone           string
	Phone_Mobile    string
	Woo_Billing     uint
	Woo_Shipping    uint
	Woo_Country     string
	Woo_State       string
	Tenant_ID       string
	Shop_ID         uint
	Foreign_ID      uint
	Customer_ID     uint
	Manufacturer_ID *uint   `gorm:"column:manufacturer_id" json:"manufacturer_id,omitempty"`
	Supplier_ID     *uint   `gorm:"column:supplier_id" json:"supplier_id,omitempty"`
	Warehouse_ID    *uint   `gorm:"column:warehouse_id" json:"warehouse_id,omitempty"`
	Country_ID      *uint   `gorm:"column:country_id" json:"country_id,omitempty"`
	State_ID        *uint   `gorm:"column:state_id" json:"state_id,omitempty"`
	DNI             *string `gorm:"column:dni" json:"dni,omitempty"`
	Alias           *string `gorm:"column:alias" json:"alias,omitempty"`
	Other           *string `gorm:"column:other" json:"other,omitempty"`
	VAT_Number      *string `gorm:"column:vat_number" json:"vat_number,omitempty"`
}

func (Addresses) TableName() string {
	return "addresses"
}
