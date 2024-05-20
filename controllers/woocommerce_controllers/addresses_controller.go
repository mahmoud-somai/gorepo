package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	AddressesRepository *repositories.AddressesRepository
}

func NewAddressRepository(addressRepository *repositories.AddressesRepository) *AddressController {
	return &AddressController{
		AddressesRepository: addressRepository,
	}
}

func (c *AddressController) CreateAddress(ctx *gin.Context) {
	var address models.Addresses

	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AddressesRepository.CreateAddress(&address); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute"})
		return
	}
	ctx.JSON(http.StatusCreated, address)
}
