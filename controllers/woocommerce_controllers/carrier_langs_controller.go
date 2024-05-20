package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CarrierLangController struct {
	CarrierLangRepository *repositories.CarrierLangRepository
}

func NewCarrierLangRepository(carrierLangRepository *repositories.CarrierLangRepository) *CarrierLangController {
	return &CarrierLangController{
		CarrierLangRepository: carrierLangRepository,
	}
}

func (c *CarrierLangController) CreateCarrierLang(ctx *gin.Context) {
	var carrier_Lang models.CarrierLang

	if err := ctx.ShouldBindJSON(&carrier_Lang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CarrierLangRepository.CreateCarrierLang(&carrier_Lang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create A Carrier Lang"})
		return
	}
	ctx.JSON(http.StatusCreated, carrier_Lang)
}
