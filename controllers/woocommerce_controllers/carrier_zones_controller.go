package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CarrierZoneController struct {
	CarrierZoneRepository *repositories.CarrierZoneRepository
}

func NewCarrierZoneRepository(carrierZoneRepository *repositories.CarrierZoneRepository) *CarrierZoneController {
	return &CarrierZoneController{
		CarrierZoneRepository: carrierZoneRepository,
	}
}

func (c *CarrierZoneController) CreateCarrierZone(ctx *gin.Context) {
	var carrier_Zone models.CarrierZone

	if err := ctx.ShouldBindJSON(&carrier_Zone); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CarrierZoneRepository.CreateCarrierZone(&carrier_Zone); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create A Carrier Zone"})
		return
	}
	ctx.JSON(http.StatusCreated, carrier_Zone)
}
