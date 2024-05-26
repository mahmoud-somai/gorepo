package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type OrderFeeController struct {
	OrderFeesRepository *repositories.OrderFeesRepository
}

func NewOrderFeeController(orderFeesRepository *repositories.OrderFeesRepository) *OrderFeeController {
	return &OrderFeeController{
		OrderFeesRepository: orderFeesRepository,
	}
}

// CreateOrderFees handles the creation of multiple order fees
func (c *OrderFeeController) CreateOrderFees(ctx *gin.Context) {
	var order_fees []models.OrderFee

	if err := ctx.ShouldBindJSON(&order_fees); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iterate over the slice and create each OrderFee
	for i, order_fee := range order_fees {
		// Get the order ID by foreign ID
		orderID, err := c.OrderFeesRepository.GetOrderIDByForeignID(int(order_fee.OrderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID", "details": err.Error()})
			return
		}

		// Set the retrieved order ID in the order fee
		order_fees[i].OrderID = int32(orderID)

		// Create the order fee
		if err := c.OrderFeesRepository.CreateOrderFee(&order_fees[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order fee", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_fees": order_fees})
}
