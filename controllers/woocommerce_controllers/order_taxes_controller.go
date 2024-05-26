package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type OrderTaxController struct {
	OrderTaxesRepository *repositories.OrderTaxesRepository
}

func NewOrderTaxController(orderTaxesRepository *repositories.OrderTaxesRepository) *OrderTaxController {
	return &OrderTaxController{
		OrderTaxesRepository: orderTaxesRepository,
	}
}

// CreateOrderTax handles the creation of multiple order taxes
func (c *OrderTaxController) CreateOrderTax(ctx *gin.Context) {
	var orderTaxes []models.OrderTax

	if err := ctx.ShouldBindJSON(&orderTaxes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range orderTaxes {
		// Get the order ID by foreign ID
		orderID, err := c.OrderTaxesRepository.GetOrderIDByForeignID(int(orderTaxes[i].OrderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID", "details": err.Error()})
			return
		}

		// Get the tax ID by foreign ID
		taxID, err := c.OrderTaxesRepository.GetTaxIDByForeignID(int(orderTaxes[i].TaxID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tax ID", "details": err.Error()})
			return
		}

		// Set the retrieved order ID and tax ID in the order tax
		orderTaxes[i].OrderID = int32(orderID)
		orderTaxes[i].TaxID = int32(taxID)

		// Create the order tax
		if err := c.OrderTaxesRepository.CreateOrderTax(&orderTaxes[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order tax", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_order_taxes": orderTaxes})
}
