package models

import (
	"time"
)

const TableNameOrder = "orders"

// Order mapped from table <orders>
type Order struct {
	ID                    int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID              string  `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	ShopID                int32   `gorm:"column:shop_id"  json:"shop_id"`
	ForeignID             int32   `gorm:"column:foreign_id" json:"foreign_id"`
	AddressDeliveryID     *int32  `gorm:"column:address_delivery_id" json:"address_delivery_id,omitempty"`
	AddressInvoiceID      *int32  `gorm:"column:address_invoice_id" json:"address_invoice_id,omitempty"`
	CartID                *int32  `gorm:"column:cart_id" json:"cart_id,omitempty"`
	CurrencyID            *int32  `gorm:"column:currency_id" json:"currency_id,omitempty"`
	LangID                int32   `gorm:"column:lang_id;not null" json:"lang_id"`
	CustomerID            int32   `gorm:"column:customer_id" json:"customer_id"`
	CarrierID             *int32  `gorm:"column:carrier_id" json:"carrier_id,omitempty"`
	Module                string  `gorm:"column:module" json:"module"`
	Reference             string  `gorm:"column:reference" json:"reference"`
	Payment               string  `gorm:"column:payment" json:"payment"`
	TotalPaidReal         float32 `gorm:"column:total_paid_real" json:"total_paid_real"`
	TotalPaid             float32 `gorm:"column:total_paid" json:"total_paid"`
	TotalPaidTaxIncl      float32 `gorm:"column:total_paid_tax_incl" json:"total_paid_tax_incl"`
	TotalPaidTaxExcl      float32 `gorm:"column:total_paid_tax_excl" json:"total_paid_tax_excl"`
	TotalProducts         float32 `gorm:"column:total_products" json:"total_products"`
	TotalProductsWt       float32 `gorm:"column:total_products_wt" json:"total_products_wt"`
	ConversionRate        float32 `gorm:"column:conversion_rate" json:"conversion_rate"`
	TotalShipping         float32 `gorm:"column:total_shipping" json:"total_shipping"`
	TotalDiscounts        float32 `gorm:"column:total_discounts" json:"total_discounts"`
	Note                  string  `gorm:"column:note" json:"note"`
	InvoiceNumber         string  `gorm:"column:invoice_number" json:"invoice_number"`
	InvoiceDate           string  `gorm:"column:invoice_date" json:"invoice_date"`
	DeliveryNumber        string  `gorm:"column:delivery_number" json:"delivery_number"`
	DeliveryDate          string  `gorm:"column:delivery_date" json:"delivery_date"`
	CarrierTaxRate        float32 `gorm:"column:carrier_tax_rate" json:"carrier_tax_rate"`
	ShippingNumber        string  `gorm:"column:shipping_number" json:"shipping_number"`
	TotalWrapping         float32 `gorm:"column:total_wrapping" json:"total_wrapping"`
	Valid                 string  `gorm:"column:valid" json:"valid"`
	CurrentState          int32   `gorm:"column:current_state;not null" json:"current_state"`
	ShopGroupIDForeign    int32   `gorm:"column:shop_group_id_foreign" json:"shop_group_id_foreign"`
	SecureKey             string  `gorm:"column:secure_key" json:"secure_key"`
	Recyclable            bool    `gorm:"column:recyclable" json:"recyclable"`
	Gift                  bool    `gorm:"column:gift" json:"gift"`
	GiftMessage           string  `gorm:"column:gift_message" json:"gift_message"`
	MobileTheme           bool    `gorm:"column:mobile_theme" json:"mobile_theme"`
	TotalDiscountsTaxIncl float32 `gorm:"column:total_discounts_tax_incl" json:"total_discounts_tax_incl"`
	TotalDiscountsTaxExcl float32 `gorm:"column:total_discounts_tax_excl" json:"total_discounts_tax_excl"`
	TotalShippingTaxIncl  float32 `gorm:"column:total_shipping_tax_incl" json:"total_shipping_tax_incl"`
	TotalShippingTaxExcl  float32 `gorm:"column:total_shipping_tax_excl" json:"total_shipping_tax_excl"`
	TotalWrappingTaxIncl  float32 `gorm:"column:total_wrapping_tax_incl" json:"total_wrapping_tax_incl"`
	TotalWrappingTaxExcl  float32 `gorm:"column:total_wrapping_tax_excl" json:"total_wrapping_tax_excl"`
	RoundMode             int32   `gorm:"column:round_mode" json:"round_mode"`
	RoundType             int32   `gorm:"column:round_type" json:"round_type"`
	DateAddForeign        string  `gorm:"column:date_add_foreign" json:"date_add_foreign"`
	WooCurrency           string  `gorm:"column:woo_currency" json:"woo_currency"`
	WooStatus             string  `gorm:"column:woo_status" json:"woo_status"`
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}

const TableNameOrderdetail = "orderdetails"

// Orderdetail mapped from table <orderdetails>
type Orderdetail struct {
	ID                        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID                  string    `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	OrderID                   int32     `gorm:"column:order_id" json:"order_id"`
	ProductID                 int32     `gorm:"column:product_id" json:"product_id"`
	ProductAttributeID        *int32    `gorm:"column:product_attribute_id" json:"product_attribute_id,omitempty"`
	ShopID                    int32     `gorm:"column:shop_id;default:1" json:"shop_id"`
	TaxRulesGroupID           *int32    `gorm:"column:tax_rules_group_id" json:"tax_rules_group_id,omitempty"`
	ForeignID                 int32     `gorm:"column:foreign_id" json:"foreign_id"`
	OrderInvoiceID            *int32    `gorm:"column:order_invoice_id" json:"order_invoice_id,omitempty"`
	IDWarehouseID             *int32    `gorm:"column:warehouse_id" json:"iwarehouse_id,omitempty"`
	CustomizationID           *int32    `gorm:"column:customization_id" json:"customization_id,omitempty"`
	ProductQuantityReinjected int32     `gorm:"column:product_quantity_reinjected" json:"product_quantity_reinjected"`
	GroupReduction            float32   `gorm:"column:group_reduction" json:"group_reduction"`
	DiscountQuantityApplied   int32     `gorm:"column:discount_quantity_applied" json:"discount_quantity_applied"`
	DownloadHash              string    `gorm:"column:download_hash" json:"download_hash"`
	DownloadDeadline          time.Time `gorm:"column:download_deadline" json:"download_deadline"`
	ProductName               string    `gorm:"column:product_name" json:"product_name"`
	ProductQuantity           int32     `gorm:"column:product_quantity" json:"product_quantity"`
	ProductQuantityInStock    int32     `gorm:"column:product_quantity_in_stock" json:"product_quantity_in_stock"`
	ProductQuantityReturn     int32     `gorm:"column:product_quantity_return" json:"product_quantity_return"`
	ProductQuantityRefunded   int32     `gorm:"column:product_quantity_refunded" json:"product_quantity_refunded"`
	ProductPrice              float32   `gorm:"column:product_price" json:"product_price"`
	ReductionPercent          float32   `gorm:"column:reduction_percent" json:"reduction_percent"`
	ReductionAmount           float32   `gorm:"column:reduction_amount" json:"reduction_amount"`
	ReductionAmountTaxIncl    float32   `gorm:"column:reduction_amount_tax_incl" json:"reduction_amount_tax_incl"`
	ReductionAmountTaxExcl    float32   `gorm:"column:reduction_amount_tax_excl" json:"reduction_amount_tax_excl"`
	ProductQuantityDiscount   float32   `gorm:"column:product_quantity_discount" json:"product_quantity_discount"`
	ProductEan13              string    `gorm:"column:product_ean13" json:"product_ean13"`
	ProductIsbn               string    `gorm:"column:product_isbn" json:"product_isbn"`
	ProductUpc                string    `gorm:"column:product_upc" json:"product_upc"`
	ProductMpn                string    `gorm:"column:product_mpn" json:"product_mpn"`
	ProductReference          string    `gorm:"column:product_reference" json:"product_reference"`
	ProductSupplierReference  string    `gorm:"column:product_supplier_reference" json:"product_supplier_reference"`
	ProductWeight             float32   `gorm:"column:product_weight" json:"product_weight"`
	TaxComputationMethod      int32     `gorm:"column:tax_computation_method" json:"tax_computation_method"`
	Ecotax                    float32   `gorm:"column:ecotax" json:"ecotax"`
	EcotaxTaxRate             float32   `gorm:"column:ecotax_tax_rate" json:"ecotax_tax_rate"`
	DownloadNb                int32     `gorm:"column:download_nb" json:"download_nb"`
	UnitPriceTaxIncl          float32   `gorm:"column:unit_price_tax_incl" json:"unit_price_tax_incl"`
	UnitPriceTaxExcl          float32   `gorm:"column:unit_price_tax_excl" json:"unit_price_tax_excl"`
	TotalPriceTaxIncl         float32   `gorm:"column:total_price_tax_incl" json:"total_price_tax_incl"`
	TotalPriceTaxExcl         float32   `gorm:"column:total_price_tax_excl" json:"total_price_tax_excl"`
	TotalShippingPriceTaxExcl float32   `gorm:"column:total_shipping_price_tax_excl" json:"total_shipping_price_tax_excl"`
	TotalShippingPriceTaxIncl float32   `gorm:"column:total_shipping_price_tax_incl" json:"total_shipping_price_tax_incl"`
	PurchaseSupplierPrice     float32   `gorm:"column:purchase_supplier_price" json:"purchase_supplier_price"`
	OriginalProductPrice      float32   `gorm:"column:original_product_price" json:"original_product_price"`
	OriginalWholesalePrice    float32   `gorm:"column:original_wholesale_price" json:"original_wholesale_price"`
	TotalRefundedTaxExcl      float32   `gorm:"column:total_refunded_tax_excl" json:"total_refunded_tax_excl"`
	TotalRefundedTaxInclt     float32   `gorm:"column:total_refunded_tax_inclt" json:"total_refunded_tax_inclt"`
}

// TableName Orderdetail's table name
func (*Orderdetail) TableName() string {
	return TableNameOrderdetail
}

const TableNameOrderCarrier = "order_carriers"

// OrderCarrier mapped from table <order_carriers>
type OrderCarrier struct {
	ID                  int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantID            string    `gorm:"column:tenant_id;default:bfb674eb-0286-4f42-8103-749a22bb8ae2" json:"tenant_id"`
	OrderID             int32     `gorm:"column:order_id" json:"order_id"`
	CarrierID           *int32    `gorm:"column:carrier_id" json:"carrier_id,omitempty"`
	ForeignID           int32     `gorm:"column:foreign_id" json:"foreign_id"`
	Name                string    `gorm:"column:name" json:"name"`
	OrderInvoice        *int32    `gorm:"column:order_invoice_id" json:"order_invoice_id,omitempty"`
	Weight              float32   `gorm:"column:weight" json:"weight"`
	ShippingCostTaxExcl float32   `gorm:"column:shipping_cost_tax_excl" json:"shipping_cost_tax_excl"`
	ShippingCostTaxIncl float32   `gorm:"column:shipping_cost_tax_incl" json:"shipping_cost_tax_incl"`
	TrackingNumber      int32     `gorm:"column:tracking_number" json:"tracking_number"`
	DateAdded           time.Time `gorm:"column:date_added" json:"date_added"`
}

// TableName OrderCarrier's table name
func (*OrderCarrier) TableName() string {
	return TableNameOrderCarrier
}

const TableNameOrderTax = "order_taxes"

// OrderTax mapped from table <order_taxes>
type OrderTax struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaxID   int32 `gorm:"column:tax_id" json:"tax_id"`
	OrderID int32 `gorm:"column:order_id" json:"order_id"`
}

// TableName OrderTax's table name
func (*OrderTax) TableName() string {
	return TableNameOrderTax
}

const TableNameOrderCarrierTax = "order_carrier_taxes"

// OrderCarrierTax mapped from table <order_carrier_taxes>
type OrderCarrierTax struct {
	ID             int32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderCarrierID int32    `gorm:"column:order_carrier_id" json:"order_carrier_id"`
	TaxID          int32    `gorm:"column:tax_id" json:"tax_id"`
	Total          float32  `gorm:"column:total" json:"total"`
	Subtotal       *float32 `gorm:"column:subtotal" json:"subtotal,omitempty"`
}

// TableName OrderCarrierTax's table name
func (*OrderCarrierTax) TableName() string {
	return TableNameOrderCarrierTax
}

const TableNameOrderDetailsTax = "order_details_taxes"

// OrderDetailsTax mapped from table <order_details_taxes>
type OrderDetailsTax struct {
	ID            int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderDetailID int32   `gorm:"column:order_detail_id" json:"order_detail_id"`
	TaxID         int32   `gorm:"column:tax_id" json:"tax_id"`
	TaxClass      string  `gorm:"column:tax_class" json:"tax_class"`
	Total         float32 `gorm:"column:total" json:"total"`
	Subtotal      float32 `gorm:"column:subtotal" json:"subtotal"`
}

// TableName OrderDetailsTax's table name
func (*OrderDetailsTax) TableName() string {
	return TableNameOrderDetailsTax
}
