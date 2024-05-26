package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	OrderDetailsRepository *repositories.OrderDetailsRepository
}

func NewOrderDetailController(orderDetailsRepository *repositories.OrderDetailsRepository) *OrderDetailController {
	return &OrderDetailController{
		OrderDetailsRepository: orderDetailsRepository,
	}
}

func (c *OrderDetailController) CreateOrderDetails(ctx *gin.Context) {
	var orderDetails []models.Orderdetail

	if err := ctx.ShouldBindJSON(&orderDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, orderDetail := range orderDetails {
		// Retrieve the product ID by the foreign ID
		productID, err := c.OrderDetailsRepository.GetProductIDByForeignID(int(orderDetail.ProductID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product ID"})
			return
		}
		orderDetail.ProductID = int32(productID)

		// Retrieve the order ID by the foreign ID
		orderID, err := c.OrderDetailsRepository.GetOrderIDByForeignID(int(orderDetail.OrderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order ID"})
			return
		}
		orderDetail.OrderID = int32(orderID)

		// Create the order detail
		if err := c.OrderDetailsRepository.CreateOrderDetail(&orderDetail); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Order Detail"})
			return
		}
	}

	ctx.JSON(http.StatusCreated, orderDetails)
}
