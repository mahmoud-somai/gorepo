package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type TaxController struct {
	TaxeRepository *repositories.TaxeRepository
}

func NewTaxController(taxeRepository *repositories.TaxeRepository) *TaxController {
	return &TaxController{
		TaxeRepository: taxeRepository,
	}
}

func (c *TaxController) CreateTaxes(ctx *gin.Context) {
	var reqs []struct {
		ForeignID uint   `json:"foreign_id"`
		Name      string `json:"name"`
		Rate      string `json:"rate"`
		WooClass  string `json:"woo_class"`
		Active    uint   `json:"active"`
		Deleted   uint   `json:"deleted"`
		LangID    uint   `json:"lang_id"`
		TenantID  string `json:"tenant_id"`
		ShopID    uint   `json:"shop_id"`
	}

	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		tax := &models.Tax{
			TenantID:  req.TenantID,
			ForeignID: req.ForeignID,
			ShopID:    req.ShopID,
			Rate:      req.Rate,
			Active:    req.Active,
			Deleted:   req.Deleted,
		}

		if err := c.TaxeRepository.CreateTaxe(tax); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Tax: " + err.Error()})
			return
		}

		taxLang := &models.TaxLang{
			TenantID: req.TenantID,
			TaxID:    tax.ID,
			LangID:   req.LangID,
			Name:     req.Name,
			WooClass: req.WooClass,
		}

		if err := c.TaxeRepository.CreateTaxeLang(taxLang); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create TaxLang: " + err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Taxes created successfully"})
}
