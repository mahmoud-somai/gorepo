package woocommerce_controllers

import (
	"log"
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDetailsTaxController struct {
	OrderDetailsTaxRepository *repositories.OrderDetailsTaxRepository
}

func NewOrderDetailsTaxController(orderDetailsTaxRepository *repositories.OrderDetailsTaxRepository) *OrderDetailsTaxController {
	return &OrderDetailsTaxController{
		OrderDetailsTaxRepository: orderDetailsTaxRepository,
	}
}

// CreateOrderDetailsTaxes handles the creation of multiple order details taxes
func (c *OrderDetailsTaxController) CreateOrderDetailsTaxes(ctx *gin.Context) {
	var orderDetailsTaxes []models.OrderDetailsTax

	if err := ctx.ShouldBindJSON(&orderDetailsTaxes); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, orderDetailsTax := range orderDetailsTaxes {
		// Get the OrderDetail ID by foreign ID
		orderDetailID, err := c.OrderDetailsTaxRepository.GetOrderDetailIDByForeignID(int(orderDetailsTax.OrderDetailID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("OrderDetail not found for foreignID %d, skipping", orderDetailsTax.OrderDetailID)
				continue
			} else {
				log.Printf("Error retrieving OrderDetail ID for foreignID %d: %v", orderDetailsTax.OrderDetailID, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order detail ID", "details": err.Error()})
				return
			}
		}

		// Get the Tax ID by foreign ID
		taxID, err := c.OrderDetailsTaxRepository.GetTaxIDByForeignID(int(orderDetailsTax.TaxID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//log.Printf("Tax not found for foreignID %d, skipping", orderDetailsTax.TaxID)
				continue
			} else {
				//log.Printf("Error retrieving Tax ID for foreignID %d: %v", orderDetailsTax.TaxID, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tax ID", "details": err.Error()})
				return
			}
		}

		// Set the retrieved IDs in the order details tax
		orderDetailsTaxes[i].OrderDetailID = int32(orderDetailID)
		orderDetailsTaxes[i].TaxID = int32(taxID)

		// Create the order details tax
		if err := c.OrderDetailsTaxRepository.CreateOrderDetailsTax(&orderDetailsTaxes[i]); err != nil {
			log.Printf("Error creating OrderDetailsTax: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order details tax", "details": err.Error()})
			return
		}
		log.Printf("Successfully created OrderDetailsTax: %v", orderDetailsTaxes[i])
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_order_details_taxes": orderDetailsTaxes})
}
