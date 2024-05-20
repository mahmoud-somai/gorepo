package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeController struct {
	AttributeRepository *repositories.AttributeRepository
}

func NewAttributeRepository(attributeRepository *repositories.AttributeRepository) *AttributeController {
	return &AttributeController{
		AttributeRepository: attributeRepository,
	}
}

func (c *AttributeController) CreateAttribute(ctx *gin.Context) {
	var attribute models.Attributes

	if err := ctx.ShouldBindJSON(&attribute); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeRepository.CreateAttribute(&attribute); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute"})
		return
	}
	ctx.JSON(http.StatusCreated, attribute)
}
