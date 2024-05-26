package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type OrderCarrierController struct {
	OrderCarriersRepository *repositories.OrderCarriersRepository
}

func NewOrderCarrierController(orderCarriersRepository *repositories.OrderCarriersRepository) *OrderCarrierController {
	return &OrderCarrierController{
		OrderCarriersRepository: orderCarriersRepository,
	}
}

// CreateOrderFees handles the creation of multiple order fees
func (c *OrderCarrierController) CreateOrderCarrier(ctx *gin.Context) {
	var order_carriers []models.OrderCarrier

	if err := ctx.ShouldBindJSON(&order_carriers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iterate over the slice and create each OrderFee
	for i, order_fee := range order_carriers {
		// Get the order ID by foreign ID
		orderID, err := c.OrderCarriersRepository.GetOrderIDByForeignID(int(order_fee.OrderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID", "details": err.Error()})
			return
		}

		// Set the retrieved order ID in the order fee
		order_carriers[i].OrderID = int32(orderID)

		// Create the order fee
		if err := c.OrderCarriersRepository.CreateOrderCarriers(&order_carriers[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order fee", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_Carrier": order_carriers})
}
