package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryShopController struct {
// 	CategoryShopRepository *repositories.CategoryShopRepository
// }

// func NewCategoryShopRepository(categoryShopRepository *repositories.CategoryShopRepository) *CategoryShopController {
// 	return &CategoryShopController{
// 		CategoryShopRepository: categoryShopRepository,
// 	}
// }

// func (c *CategoryShopController) CreateCategoryShop(ctx *gin.Context) {
// 	var category_Shop models.CategoryShop

// 	if err := ctx.ShouldBindJSON(&category_Shop); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CategoryShopRepository.CreateCategoryShop(&category_Shop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Category Shop"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, category_Shop)
// }
