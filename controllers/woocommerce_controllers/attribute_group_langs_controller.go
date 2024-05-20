package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeGroupLangController struct {
	AttributeGroupLangRepository *repositories.AttributeGroupLangRepository
}

func NewAttributeGroupLangRepository(attributeGroupLangRepository *repositories.AttributeGroupLangRepository) *AttributeGroupLangController {
	return &AttributeGroupLangController{
		AttributeGroupLangRepository: attributeGroupLangRepository,
	}
}

func (c *AttributeGroupLangController) CreateAttributeGroupLang(ctx *gin.Context) {
	var attributeGroupLang models.AttributeGroupLang

	if err := ctx.ShouldBindJSON(&attributeGroupLang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeGroupLangRepository.CreateAttributeGroupLang(&attributeGroupLang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute Group Lang"})
		return
	}
	ctx.JSON(http.StatusCreated, attributeGroupLang)
}
