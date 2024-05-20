package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeShopController struct {
	AttributeShopRepository *repositories.AttributeShopRepository
}

func NewAttributeShopRepository(attributeShopRepository *repositories.AttributeShopRepository) *AttributeShopController {
	return &AttributeShopController{
		AttributeShopRepository: attributeShopRepository,
	}
}

func (c *AttributeShopController) CreateAttributeShop(ctx *gin.Context) {
	var attributeShop models.AttributeShop

	if err := ctx.ShouldBindJSON(&attributeShop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeShopRepository.CreateAttributeShop(&attributeShop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute Shop"})
		return
	}
	ctx.JSON(http.StatusCreated, attributeShop)
}
