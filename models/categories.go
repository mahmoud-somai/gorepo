package models

type Category struct {
	ID               uint `gorm:"primaryKey"`
	Tenant_ID        string
	Parent_ID        uint
	Level_Depth      *int    `gorm:"column:level_depth" json:"level_depth,omitempty"`
	Active           *uint   `gorm:"column:active" json:"active,omitempty"`
	Is_Root_Category *uint   `gorm:"column:is_root_category" json:"is_root_category,omitempty"`
	Position         *int    `gorm:"column:position" json:"position,omitempty"`
	Meta_Title       *string `gorm:"column:meta_title" json:"meta_title,omitempty"`
	Meta_Description *string `gorm:"column:meta_description" json:"meta_description,omitempty"`
	Meta_Keywords    *string `gorm:"column:meta_keywords" json:"meta_keywords,omitempty"`
	Link_Rewrite     *string `gorm:"column:link_rewrite" json:"link_rewrite,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryLang struct {
	ID          uint `gorm:"primaryKey"`
	Tenant_ID   string
	Category_ID uint
	Lang_ID     uint
	Name        string
	Description string
}

func (CategoryLang) TableName() string {
	return "category_langs"
}

type CategoryProduct struct {
	ID          uint `gorm:"primaryKey"`
	Tenant_ID   string
	Category_ID uint
	Product_ID  uint
}

func (CategoryProduct) TableName() string {
	return "category_product"
}

type CategoryShop struct {
	ID          uint `gorm:"primaryKey"`
	Tenant_ID   string
	Category_ID uint
	Shop_ID     uint
	Foreign_ID  uint
}

func (CategoryShop) TableName() string {
	return "category_shops"
}
