package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type FeatureController struct {
	FeatureRepository *repositories.FeatureRepository
}

func NewFeatureRepository(featureRepository *repositories.FeatureRepository) *FeatureController {
	return &FeatureController{
		FeatureRepository: featureRepository,
	}
}

func (c *FeatureController) CreateFeature(ctx *gin.Context) {
	var feature models.Feature

	if err := ctx.ShouldBindJSON(&feature); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.FeatureRepository.CreateFeature(&feature); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Feature"})
		return
	}
	ctx.JSON(http.StatusCreated, feature)
}

func (c *FeatureController) CreateFeatureLang(ctx *gin.Context) {
	var feature_lang models.FeatureLang

	if err := ctx.ShouldBindJSON(&feature_lang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.FeatureRepository.CreateFeatureLang(&feature_lang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Feature Lang"})
		return
	}
	ctx.JSON(http.StatusCreated, feature_lang)
}

func (c *FeatureController) CreateFeatureShop(ctx *gin.Context) {
	var feature_shop models.FeatureShop

	if err := ctx.ShouldBindJSON(&feature_shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.FeatureRepository.CreateFeatureShop(&feature_shop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Feature Shop"})
		return
	}
	ctx.JSON(http.StatusCreated, feature_shop)
}

func (c *FeatureController) CreateFeatureValue(ctx *gin.Context) {
	var feature_value models.FeatureValue

	if err := ctx.ShouldBindJSON(&feature_value); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.FeatureRepository.CreateFeatureValue(&feature_value); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Feature Value"})
		return
	}
	ctx.JSON(http.StatusCreated, feature_value)
}

func (c *FeatureController) CreateFeatureValueLang(ctx *gin.Context) {
	var feature_value_lang models.FeatureValueLang

	if err := ctx.ShouldBindJSON(&feature_value_lang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.FeatureRepository.CreateFeatureValueLang(&feature_value_lang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Feature Value Lang"})
		return
	}
	ctx.JSON(http.StatusCreated, feature_value_lang)
}
