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

func NewBillingRepository(billingRepository *repositories.BillingRepository) *BillingController {
	return &BillingController{
		BillingRepository: billingRepository,
	}
}

func (c *BillingController) CreateBilling(ctx *gin.Context) {
	var billing models.Billing_addresses_woo

	if err := ctx.ShouldBindJSON(&billing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.BillingRepository.CreateBilling(&billing); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create A Billing"})
		return
	}
	ctx.JSON(http.StatusCreated, billing)
}
