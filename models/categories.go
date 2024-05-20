package models

type Category struct {
	ID               uint `gorm:"primaryKey"`
	Tenant_ID        string
	Parent_ID        uint
	Level_Depth      int
	Active           uint
	Is_Root_Category uint
	Position         int
	Meta_Title       string
	Meta_Description string
	Meta_Keywords    string
	Link_Rewrite     string
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
