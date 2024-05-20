package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeGroupController struct {
	AttributeGroupRepository *repositories.AttributeGroupRepository
}

func NewAttributeGroupRepository(attributeGroupRepository *repositories.AttributeGroupRepository) *AttributeGroupController {
	return &AttributeGroupController{
		AttributeGroupRepository: attributeGroupRepository,
	}
}

func (c *AttributeGroupController) CreateAttributeGroup(ctx *gin.Context) {
	var attributeGroup models.AttributeGroup

	if err := ctx.ShouldBindJSON(&attributeGroup); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeGroupRepository.CreateAttributeGroup(&attributeGroup); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute Group"})
		return
	}
	ctx.JSON(http.StatusCreated, attributeGroup)
}
