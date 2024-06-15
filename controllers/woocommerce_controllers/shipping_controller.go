package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingController struct {
	ShippingRepository *repositories.ShippingRepository
}

func NewShippingController(shippingRepository *repositories.ShippingRepository) *ShippingController {
	return &ShippingController{
		ShippingRepository: shippingRepository,
	}
}

// CreateShipping handles the creation of shipping addresses
func (c *ShippingController) CreateShipping(ctx *gin.Context) {
	var shippingAddresses []models.ShippingAddressesWoo

	if err := ctx.ShouldBindJSON(&shippingAddresses); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range shippingAddresses {
		// Get the order ID by foreign ID
		orderID, err := c.ShippingRepository.GetOrderIDByForeignID(int(shippingAddresses[i].OrderID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Skip the order detail if the order is not found
				continue
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order ID"})
				return
			}
		}

		// Set the retrieved order ID in the shipping address
		shippingAddresses[i].OrderID = int32(orderID)

		// Create the shipping address
		if err := c.ShippingRepository.CreateShipping(&shippingAddresses[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shipping address", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_shipping_addresses": shippingAddresses})
}
