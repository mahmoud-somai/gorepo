package models

const TableNameShop = "shops"

// Shop mapped from table <shops>
type Shop struct {
	ID                  int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID            string `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	GroupShopID         int32  `gorm:"column:group_shop_id;not null" json:"group_shop_id"`
	Name                string `gorm:"column:name" json:"name"`
	Type                string `gorm:"column:type" json:"type"`
	URL                 string `gorm:"column:url" json:"url"`
	PresKey             string `gorm:"column:pres_key" json:"pres_key"`
	WooConsumerKey      string `gorm:"column:woo_consumer_key" json:"woo_consumer_key"`
	WooConsumerSecret   string `gorm:"column:woo_consumer_secret" json:"woo_consumer_secret"`
	WooVersion          string `gorm:"column:woo_version" json:"woo_version"`
	Active              bool   `gorm:"column:active" json:"active"`
	IsImported          bool   `gorm:"column:is_imported" json:"is_imported"`
	IsEmployeesImported bool   `gorm:"column:is_employees_imported" json:"is_employees_imported"`
	UserPlanID          int32  `gorm:"column:user_plan_id" json:"user_plan_id"`
	Default             bool   `gorm:"column:default" json:"default"`
}

// TableName Shop's table name
func (*Shop) TableName() string {
	return TableNameShop
}
