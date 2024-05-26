package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type BillingController struct {
	BillingRepository *repositories.BillingRepository
}

func NewBillingController(billingRepository *repositories.BillingRepository) *BillingController {
	return &BillingController{
		BillingRepository: billingRepository,
	}
}

// CreateBilling handles the creation of billing addresses
func (c *BillingController) CreateBilling(ctx *gin.Context) {
	var billingAddresses []models.BillingAddressesWoo

	if err := ctx.ShouldBindJSON(&billingAddresses); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range billingAddresses {
		// Get the order ID by foreign ID
		orderID, err := c.BillingRepository.GetOrderIDByForeignID(int(billingAddresses[i].OrderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID", "details": err.Error()})
			return
		}

		// Set the retrieved order ID in the billing address
		billingAddresses[i].OrderID = int32(orderID)

		// Create the billing address
		if err := c.BillingRepository.CreateBilling(&billingAddresses[i]); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create billing address", "details": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "created_billing_addresses": billingAddresses})
}
