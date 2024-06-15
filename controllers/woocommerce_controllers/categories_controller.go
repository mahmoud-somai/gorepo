package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryRepository *repositories.CategoryRepository
}

func NewCategoryController(categoryRepository *repositories.CategoryRepository) *CategoryController {
	return &CategoryController{
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryController) CreateFullCategory(ctx *gin.Context) {
	var reqs []struct {
		TenantID        string  `json:"tenant_id"`
		ParentID        uint    `json:"parent"`
		ForeignID       uint    `json:"foreign_id"`
		Name            string  `json:"name"`
		Description     string  `json:"description"`
		LevelDepth      *int    `json:"level_depth"`
		Active          *uint   `json:"active"`
		IsRootCategory  *uint   `json:"is_root_category"`
		Position        *int    `json:"position"`
		MetaTitle       *string `json:"meta_title"`
		MetaDescription *string `json:"meta_description"`
		MetaKeywords    *string `json:"meta_keywords"`
		LinkRewrite     *string `json:"link_rewrite"`
		LangID          uint    `json:"lang_id"`
		ShopID          uint    `json:"shop_id"`
	}

	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iterate over each request and process the creation
	for _, req := range reqs {
		// Check if the category already exists
		exists, err := c.CategoryRepository.CategoryExists(req.TenantID, req.ForeignID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if category exists"})
			return
		}

		// Skip creation if category already exists
		if exists {
			continue
		}

		category := &models.Category{
			Tenant_ID: req.TenantID,
			Parent_ID: req.ParentID,
		}

		categoryLang := &models.CategoryLang{
			Tenant_ID:   req.TenantID,
			Lang_ID:     req.LangID,
			Name:        req.Name,
			Description: req.Description,
		}

		categoryShop := &models.CategoryShop{
			Tenant_ID:  req.TenantID,
			Shop_ID:    req.ShopID,
			Foreign_ID: req.ForeignID,
		}

		// Use the repository method to create the full category
		if err := c.CategoryRepository.CreateFullCategory(category, categoryLang, categoryShop); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Categories created successfully"})
}
