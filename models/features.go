package models

type Feature struct {
	ID        uint `gorm:"primaryKey"`
	Tenant_ID string
	Position  uint
}

func (Feature) TableName() string {
	return "features"
}

type FeatureLang struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Lang_ID    uint
	Feature_ID uint
	Name       string
}

func (FeatureLang) TableName() string {
	return "feature_langs"
}

type FeatureShop struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Feature_ID uint
	Shop_ID    uint
	Foreign_ID uint
}

func (FeatureShop) TableName() string {
	return "feature_shops"
}

type FeatureValue struct {
	ID         uint `gorm:"primaryKey"`
	Tenant_ID  string
	Feature_ID uint
	Custom     string
}

func (FeatureValue) TableName() string {
	return "feature_values"
}

type FeatureValueLang struct {
	ID               uint `gorm:"primaryKey"`
	Tenant_ID        string
	Feature_Value_ID uint
	Lang_ID          uint
	Value            string
}

func (FeatureValueLang) TableName() string {
	return "feature_value_langs"
}
