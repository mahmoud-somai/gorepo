package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CarrierShopController struct {
	CarrierShopRepository *repositories.CarrierShopRepository
}

func NewCarrierShopRepository(carrierShopRepository *repositories.CarrierShopRepository) *CarrierShopController {
	return &CarrierShopController{
		CarrierShopRepository: carrierShopRepository,
	}
}

func (c *CarrierShopController) CreateCarrierShop(ctx *gin.Context) {
	var carrier_Shop models.CarrierShop

	if err := ctx.ShouldBindJSON(&carrier_Shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CarrierShopRepository.CreateCarrierShop(&carrier_Shop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create A Carrier Shop"})
		return
	}
	ctx.JSON(http.StatusCreated, carrier_Shop)
}
