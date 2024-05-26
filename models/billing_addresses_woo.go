package models

import (
	"time"
)

const TableNameBillingAddressesWoo = "billing_addresses_woo"

// BillingAddressesWoo mapped from table <billing_addresses_woo>
type BillingAddressesWoo struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID   int32     `gorm:"column:order_id" json:"order_id"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Company   string    `gorm:"column:company" json:"company"`
	Address1  string    `gorm:"column:address_1" json:"address_1"`
	Address2  string    `gorm:"column:address_2" json:"address_2"`
	City      string    `gorm:"column:city" json:"city"`
	State     string    `gorm:"column:state" json:"state"`
	Postcode  string    `gorm:"column:postcode" json:"postcode"`
	Country   string    `gorm:"column:country" json:"country"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:createdAt;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null" json:"updatedAt"`
}

// TableName BillingAddressesWoo's table name
func (*BillingAddressesWoo) TableName() string {
	return TableNameBillingAddressesWoo
}
