package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type DeliveryController struct {
	DeliveryRepository *repositories.DeliveryRepository
}

func NewDeliveryRepository(deliveryRepository *repositories.DeliveryRepository) *DeliveryController {
	return &DeliveryController{
		DeliveryRepository: deliveryRepository,
	}
}

func (c *DeliveryController) CreateDelivery(ctx *gin.Context) {
	var delivery models.Delivery

	if err := ctx.ShouldBindJSON(&delivery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.DeliveryRepository.CreateDelivery(&delivery); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Delivery"})
		return
	}
	ctx.JSON(http.StatusCreated, delivery)
}
