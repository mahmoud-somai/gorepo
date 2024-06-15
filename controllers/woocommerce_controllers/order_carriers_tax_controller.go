package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderCarrierTaxController struct {
	OrderCarrierTaxRepository *repositories.OrderCarrierTaxRepository
}

func NewOrderCarrierTaxController(orderCarrierTaxRepository *repositories.OrderCarrierTaxRepository) *OrderCarrierTaxController {
	return &OrderCarrierTaxController{
		OrderCarrierTaxRepository: orderCarrierTaxRepository,
	}
}

// CreateOrderCarrierTaxes handles the creation of multiple order carrier taxes
func (c *OrderCarrierTaxController) CreateOrderCarrierTaxes(ctx *gin.Context) {
	var orderCarrierTaxes []models.OrderCarrierTax

	if err := ctx.ShouldBindJSON(&orderCarrierTaxes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, orderCarrierTax := range orderCarrierTaxes {
		// Get the OrderCarrier ID by foreign ID
		orderCarrierID, err := c.OrderCarrierTaxRepository.GetOrderCarrierIDByForeignID(int(orderCarrierTax.OrderCarrierID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//log.Printf("Order carrier not found for foreign ID: %d", orderCarrierTax.OrderCarrierID)
				continue
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order carrier ID", "details": err.Error()})
				return
			}
		}

		// Get the Tax ID by foreign ID
		taxID, err := c.OrderCarrierTaxRepository.GetTaxIDByForeignID(int(orderCarrierTax.TaxID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//log.Printf("Tax not found for foreign ID: %d", orderCarrierTax.TaxID)
				continue
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tax ID", "details": err.Error()})
				return
			}
		}

		// Set the retrieved IDs in the order carrier tax
		orderCarrierTaxes[i].OrderCarrierID = int32(orderCarrierID)
		orderCarrierTaxes[i].TaxID = int32(taxID)

		// Create the order carrier tax
		if err := c.OrderCarrierTaxRepository.CreateOrderCarrierTax(&orderCarrierTaxes[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order carrier tax", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_order_carrier_taxes": orderCarrierTaxes})
}
