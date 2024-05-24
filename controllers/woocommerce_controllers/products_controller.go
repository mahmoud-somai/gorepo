package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductRepository *repositories.ProductRepository
}

func NewProductController(productRepository *repositories.ProductRepository) *ProductController {
	return &ProductController{
		ProductRepository: productRepository,
	}
}

func (pc *ProductController) CreateFullProducts(ctx *gin.Context) {
	var reqs []struct {
		TenantID               string  `json:"tenant_id"`
		ForeignID              uint    `json:"foreign_id"`
		ShopID                 uint    `json:"shop_id"`
		LangID                 uint    `json:"lang_id"`
		TaxID                  uint    `json:"tax_id"`
		Ecotax                 float64 `json:"ecotax"`
		DefaultImageID         uint    `json:"default_image_id"`
		ManufacturerID         *uint   `json:"manufacturer_id"` // Pointer to allow NULL
		Name                   string  `json:"name"`
		Reference              string  `json:"reference"`
		Slug                   string  `json:"slug"`
		Type                   string  `json:"type"`
		Status                 string  `json:"status"`
		Description            string  `json:"description"`
		ShortDescription       string  `json:"short_description"`
		Price                  float64 `json:"price"`
		RegularPrice           float64 `json:"regular_price"`
		DateOnSaleFrom         string  `json:"date_on_sale_from"`
		DateOnSaleTo           string  `json:"date_on_sale_to"`
		OnSale                 bool    `json:"on_sale"`
		Purchasable            bool    `json:"purchasable"`
		Height                 float64 `json:"height"`
		Width                  float64 `json:"width"`
		Length                 float64 `json:"length"`
		Weight                 float64 `json:"weight"`
		EAN13                  string  `json:"ean13"`
		ISBN                   string  `json:"isbn"`
		UPC                    string  `json:"upc"`
		MPN                    string  `json:"mpn"`
		MetaTitle              string  `json:"meta_title"`
		MetaDescription        string  `json:"meta_description"`
		MetaKeywords           string  `json:"meta_keywords"`
		Unity                  string  `json:"unity"`
		UnitPriceRatio         float64 `json:"unit_price_ratio"`
		AdditionalShippingCost float64 `json:"additional_shipping_cost"`
		File                   string  `json:"file"`
		Condition              string  `json:"condition"`
		Active                 bool    `json:"active"`
	}

	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		product := &models.Product{
			TenantID:         req.TenantID,
			ProductID:        req.ForeignID,
			DefaultImageID:   req.DefaultImageID,
			ManufacturerID:   req.ManufacturerID, // Nullable field
			TaxRuleGroupID:   req.TaxID,
			Name:             req.Name,
			Reference:        req.Reference,
			Slug:             req.Slug,
			Type:             req.Type,
			Status:           req.Status,
			Description:      req.Description,
			ShortDescription: req.ShortDescription,
			Ean13:            req.EAN13,
			Isbn:             req.ISBN,
			Upc:              req.UPC,
			Mpn:              req.MPN,
		}

		productLang := &models.ProductLang{
			TenantID:         req.TenantID,
			LangID:           req.LangID,
			ShopID:           req.ShopID,
			ProductID:        req.ForeignID,
			Name:             req.Name,
			Description:      req.Description,
			ShortDescription: req.ShortDescription,
			MetaTitle:        req.MetaTitle,
			MetaDescription:  req.MetaDescription,
			MetaKeywords:     req.MetaKeywords,
		}

		productShop := &models.ProductShop{
			TenantID:               req.TenantID,
			ShopID:                 req.ShopID,
			ProductID:              req.ForeignID,
			TaxID:                  req.TaxID,
			ForeignID:              req.ForeignID,
			Ecotax:                 float32(req.Ecotax),
			Price:                  float32(req.Price),
			Unity:                  req.Unity,
			UnityPriceRatio:        float32(req.UnitPriceRatio),
			AdditionalShippingCost: float32(req.AdditionalShippingCost),
			File:                   req.File,
			Condition:              req.Condition,
			Active:                 req.Active,
		}

		if err := pc.ProductRepository.CreateFullProduct(product, productLang, productShop); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Product"})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Products created successfully"})
}
