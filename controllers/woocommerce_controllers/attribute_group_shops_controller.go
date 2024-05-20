package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeGroupShopController struct {
	AttributeGroupShopRepository *repositories.AttributeGroupShopRepository
}

func NewAttributeGroupShopRepository(attributeGroupShopRepository *repositories.AttributeGroupShopRepository) *AttributeGroupShopController {
	return &AttributeGroupShopController{
		AttributeGroupShopRepository: attributeGroupShopRepository,
	}
}

func (c *AttributeGroupShopController) CreateAttributeGroupShop(ctx *gin.Context) {
	var attributeGroupShop models.AttributeGroupShop

	if err := ctx.ShouldBindJSON(&attributeGroupShop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeGroupShopRepository.CreateAttributeGroupShop(&attributeGroupShop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute Group Shop"})
		return
	}
	ctx.JSON(http.StatusCreated, attributeGroupShop)
}
