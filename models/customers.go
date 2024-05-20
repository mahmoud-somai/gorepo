// customer.go

package models

import "time"

type Customer struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Foreign_ID uint
	Shop_ID    uint
	First_Name string
	Last_Name  string
	Email      string
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerMessage struct {
	ID                 uint `gorm:"primaryKey"`
	User_ID            uint
	Customer_Thread_ID uint
	IP_Address         string
	Message            string
	File_Name          string
	User_Agent         string
	Private            uint
	Date_Add           time.Time
	Read               uint
}

func (CustomerMessage) TableName() string {
	return "customer_messages"
}

type CustomerShop struct {
	ID          uint `gorm:"primaryKey"`
	Tenant_ID   string
	Shop_ID     uint
	Customer_ID uint
	Foreign_ID  uint
}

func (CustomerShop) TableName() string {
	return "customer_shops"
}

type CustomerThread struct {
	ID          uint `gorm:"primaryKey"`
	Foreign_ID  uint
	Lang_ID     uint
	Shop_ID     uint
	Customer_ID uint
	Contact_ID  uint
	Order_ID    uint
	Email       string
	Token       string
	Status      uint
	Date_Add    time.Time
}

func (CustomerThread) TableName() string {
	return "customer_threads"
}
