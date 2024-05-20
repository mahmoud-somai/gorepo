package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CountryController struct {
	CountryRepository *repositories.CountryRepository
}

func NewCountryRepository(countryRepository *repositories.CountryRepository) *CountryController {
	return &CountryController{
		CountryRepository: countryRepository,
	}
}

func (c *CountryController) CreateCountry(ctx *gin.Context) {
	var country models.Country

	if err := ctx.ShouldBindJSON(&country); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CountryRepository.CreateCountry(&country); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Country"})
		return
	}
	ctx.JSON(http.StatusCreated, country)
}

func (c *CountryController) CreateCountryLang(ctx *gin.Context) {
	var country_lang models.CountryLang

	if err := ctx.ShouldBindJSON(&country_lang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CountryRepository.CreateCountryLang(&country_lang); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Country Langs"})
		return
	}
	ctx.JSON(http.StatusCreated, country_lang)
}

func (c *CountryController) CreateCountryShop(ctx *gin.Context) {
	var country_shop models.CountryShop

	if err := ctx.ShouldBindJSON(&country_shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You might want to validate the example data here if needed

	if err := c.CountryRepository.CreateCountryShop(&country_shop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Country shops"})
		return
	}
	ctx.JSON(http.StatusCreated, country_shop)
}
