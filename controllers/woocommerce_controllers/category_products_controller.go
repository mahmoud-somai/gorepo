package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryProductController struct {
// 	CategoryProductRepository *repositories.CategoryProductRepository
// }

// func NewCategoryProductRepository(categoryProductRepository *repositories.CategoryProductRepository) *CategoryProductController {
// 	return &CategoryProductController{
// 		CategoryProductRepository: categoryProductRepository,
// 	}
// }

// func (c *CategoryProductController) CreateCategoryProduct(ctx *gin.Context) {
// 	var category_Product models.CategoryProduct

// 	if err := ctx.ShouldBindJSON(&category_Product); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CategoryProductRepository.CreateCategoryProduct(&category_Product); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create CategoryProduct"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, category_Product)
// }
