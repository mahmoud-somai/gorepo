package models

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID               uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID         string    `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	ProductID        uint      `gorm:"column:product_id" json:"product_id"`
	DefaultImageID   uint      `gorm:"column:default_image_id" json:"default_image_id"`
	ManufacturerID   *uint     `gorm:"column:manufacturer_id" json:"manufacturer_id"` // Pointer to uint to allow NULL
	TaxRuleGroupID   uint      `gorm:"column:tax_rule_group_id" json:"tax_rule_group_id"`
	Name             string    `gorm:"column:name" json:"name"`
	Reference        string    `gorm:"column:reference" json:"reference"`
	Slug             string    `gorm:"column:slug" json:"slug"`
	Type             string    `gorm:"column:type" json:"type"`
	Status           string    `gorm:"column:status;not null" json:"status"`
	Description      string    `gorm:"column:description" json:"description"`
	ShortDescription string    `gorm:"column:short_description" json:"short_description"`
	Price            string    `gorm:"column:price" json:"price"`
	RegularPrice     string    `gorm:"column:regular_price" json:"regular_price"`
	DateOnSaleFrom   time.Time `gorm:"column:date_on_sale_from" json:"date_on_sale_from"`
	DateOnSaleTo     time.Time `gorm:"column:date_on_sale_to" json:"date_on_sale_to"`
	OnSale           bool      `gorm:"column:on_sale" json:"on_sale"`
	Purchasable      bool      `gorm:"column:purchasable;default:1" json:"purchasable"`
	Height           string    `gorm:"column:height" json:"height"`
	Width            string    `gorm:"column:width" json:"width"`
	Length           string    `gorm:"column:length" json:"length"`
	Weight           string    `gorm:"column:weight" json:"weight"`
	Ean13            string    `gorm:"column:ean13" json:"ean13"`
	Isbn             string    `gorm:"column:isbn" json:"isbn"`
	Upc              string    `gorm:"column:upc" json:"upc"`
	Mpn              string    `gorm:"column:mpn" json:"mpn"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}

const TableNameProductShop = "product_shops"

// ProductShop mapped from table <product_shops>
type ProductShop struct {
	ID                     uint    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID               string  `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	ProductID              uint    `gorm:"column:product_id;not null" json:"product_id"`
	ShopID                 uint    `gorm:"column:shop_id;not null" json:"shop_id"`
	TaxID                  uint    `gorm:"column:tax_id" json:"tax_id"`
	ForeignID              uint    `gorm:"column:foreign_id;not null" json:"foreign_id"`
	Ecotax                 float32 `gorm:"column:ecotax" json:"ecotax"`
	Price                  float32 `gorm:"column:price" json:"price"`
	Unity                  string  `gorm:"column:unity" json:"unity"`
	UnityPriceRatio        float32 `gorm:"column:unity_price_ratio" json:"unity_price_ratio"`
	AdditionalShippingCost float32 `gorm:"column:additional_shipping_cost" json:"additional_shipping_cost"`
	File                   string  `gorm:"column:file" json:"file"`
	Condition              string  `gorm:"column:condition" json:"condition"`
	Active                 bool    `gorm:"column:active" json:"active"`
}

// TableName ProductShop's table name
func (*ProductShop) TableName() string {
	return TableNameProductShop
}

const TableNameProductLang = "product_langs"

// ProductLang mapped from table <product_langs>
type ProductLang struct {
	ID               uint   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID         string `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	ProductID        uint   `gorm:"column:product_id;not null" json:"product_id"`
	LangID           uint   `gorm:"column:lang_id;not null" json:"lang_id"`
	ShopID           uint   `gorm:"column:shop_id;not null" json:"shop_id"`
	Name             string `gorm:"column:name" json:"name"`
	Description      string `gorm:"column:description" json:"description"`
	ShortDescription string `gorm:"column:short_description" json:"short_description"`
	MetaTitle        string `gorm:"column:meta_title" json:"meta_title"`
	MetaDescription  string `gorm:"column:meta_description" json:"meta_description"`
	MetaKeywords     string `gorm:"column:meta_keywords" json:"meta_keywords"`
}

// TableName ProductLang's table name
// func (*ProductLang) TableName() string {
// 	return TableNameProductLang
// }

// const TableNameProductAttribute = "product_attributes"

// // ProductAttribute mapped from table <product_attributes>
// type ProductAttribute struct {
// 	ID        int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
// 	TenantID  string  `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
// 	ProductID int32   `gorm:"column:product_id;not null" json:"product_id"`
// 	Reference string  `gorm:"column:reference" json:"reference"`
// 	Ean       string  `gorm:"column:ean" json:"ean"`
// 	Isbn      string  `gorm:"column:isbn" json:"isbn"`
// 	Upc       string  `gorm:"column:upc" json:"upc"`
// 	Mpn       string  `gorm:"column:mpn" json:"mpn"`
// 	Price     float32 `gorm:"column:price" json:"price"`
// 	Ecotax    float32 `gorm:"column:ecotax" json:"ecotax"`
// 	Quantity  int32   `gorm:"column:quantity" json:"quantity"`
// 	Width     float32 `gorm:"column:width" json:"width"`
// 	Height    float32 `gorm:"column:height" json:"height"`
// 	Depth     float32 `gorm:"column:depth" json:"depth"`
// 	Weight    float32 `gorm:"column:weight" json:"weight"`
// 	DefaultOn bool    `gorm:"column:default_on" json:"default_on"`
// }

// // TableName ProductAttribute's table name
// func (*ProductAttribute) TableName() string {
// 	return TableNameProductAttribute
// }

// const TableNameProductAttributeVariation = "product_attribute_variations"

// // ProductAttributeVariation mapped from table <product_attribute_variations>
// type ProductAttributeVariation struct {
// 	ID                 int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
// 	TenantID           string `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
// 	AttributeID        int32  `gorm:"column:attribute_id;not null" json:"attribute_id"`
// 	ProductAttributeID int32  `gorm:"column:product_attribute_id;not null" json:"product_attribute_id"`
// }

// // TableName ProductAttributeVariation's table name
// func (*ProductAttributeVariation) TableName() string {
// 	return TableNameProductAttributeVariation
// }

// const TableNameProductAttributeShop = "product_attribute_shops"

// // ProductAttributeShop mapped from table <product_attribute_shops>
// type ProductAttributeShop struct {
// 	ID        uint    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
// 	TenantID  string  `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
// 	ProductID uint    `gorm:"column:product_id;not null" json:"product_id"`
// 	ShopID    uint    `gorm:"column:shop_id;not null" json:"shop_id"`
// 	Price     float32 `gorm:"column:price" json:"price"`
// 	Ecotax    float32 `gorm:"column:ecotax" json:"ecotax"`
// }

// // TableName ProductAttributeShop's table name
// func (*ProductAttributeShop) TableName() string {
// 	return TableNameProductAttributeShop
// }
