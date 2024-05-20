package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CarrierController struct {
	CarrierRepository *repositories.CarrierRepository
}

func NewCarrierRepository(carrierRepository *repositories.CarrierRepository) *CarrierController {
	return &CarrierController{
		CarrierRepository: carrierRepository,
	}
}

func (c *CarrierController) CreateCarrier(ctx *gin.Context) {
	var carrier models.Carrier

	if err := ctx.ShouldBindJSON(&carrier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CarrierRepository.CreateCarrier(&carrier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create A Carrier"})
		return
	}
	ctx.JSON(http.StatusCreated, carrier)
}
