package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type AttributeLangController struct {
	AttributeLangRepository *repositories.AttributeLangRepository
}

func NewAttributeLangRepository(attributeLangRepository *repositories.AttributeLangRepository) *AttributeLangController {
	return &AttributeLangController{
		AttributeLangRepository: attributeLangRepository,
	}
}

func (c *AttributeLangController) CreateAttributeLang(ctx *gin.Context) {
	var attributeLang models.AttributeLang

	if err := ctx.ShouldBindJSON(&attributeLang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.AttributeLangRepository.CreateAttributeLang(&attributeLang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Attribute  Lang"})
		return
	}
	ctx.JSON(http.StatusCreated, attributeLang)
}
