package models

const TableNameOrderFee = "order_fees"

// OrderFee mapped from table <order_fees>
type OrderFee struct {
	ID        int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID   int32   `gorm:"column:order_id" json:"order_id"`
	ForeignID int32   `gorm:"column:foreign_id" json:"foreign_id"`
	Name      string  `gorm:"column:name" json:"name"`
	TaxClass  string  `gorm:"column:tax_class" json:"tax_class"`
	Total     float32 `gorm:"column:total" json:"total"`
	TotalTax  float32 `gorm:"column:total_tax" json:"total_tax"`
}

// TableName OrderFee's table name
func (*OrderFee) TableName() string {
	return TableNameOrderFee
}
