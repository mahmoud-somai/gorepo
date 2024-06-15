package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
	Repo *repositories.ShopRepository
}

func NewShopController(repo *repositories.ShopRepository) *ShopController {
	return &ShopController{Repo: repo}
}

func (c *ShopController) GetShopByURL(ctx *gin.Context) {
	url := ctx.Query("url") // Use query parameter instead of URL parameter
	shop, err := c.Repo.GetShopByURL(url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"shop_id":   shop.ID,
		"tenant_id": shop.TenantID,
	})
}
