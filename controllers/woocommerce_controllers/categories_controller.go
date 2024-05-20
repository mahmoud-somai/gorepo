// package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryController struct {
// 	CategoryRepository *repositories.CategoryRepository
// }

// func NewCategoryRepository(categoryRepository *repositories.CategoryRepository) *CategoryController {
// 	return &CategoryController{
// 		CategoryRepository: categoryRepository,
// 	}
// }

// func (c *CategoryController) CreateCategory(ctx *gin.Context) {
// 	var categories []models.Category

// 	if err := ctx.ShouldBindJSON(&categories); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the data here if needed

// 	if err := c.CategoryRepository.CreateCategories(categories); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Categories"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Categories created successfully"})
// }

// func (c *CategoryController) CreateCategoryLang(ctx *gin.Context) {
// 	var category_Lang models.CategoryLang

// 	if err := ctx.ShouldBindJSON(&category_Lang); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CategoryRepository.CreateCategoryLang(&category_Lang); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create CategoryLang"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, category_Lang)
// }

// func (c *CategoryController) CreateCategoryProduct(ctx *gin.Context) {
// 	var category_Product models.CategoryProduct

// 	if err := ctx.ShouldBindJSON(&category_Product); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CategoryRepository.CreateCategoryProduct(&category_Product); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create CategoryProduct"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, category_Product)
// }

// func (c *CategoryController) CreateCategoryShop(ctx *gin.Context) {
// 	var category_Shop models.CategoryShop

// 	if err := ctx.ShouldBindJSON(&category_Shop); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CategoryRepository.CreateCategoryShop(&category_Shop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Category Shop"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, category_Shop)
// }

// // func (c *CategoryController) CreateCategory(ctx *gin.Context) {
// // 	var category models.Category

// // 	if err := ctx.ShouldBindJSON(&category); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	// You might want to validate the example data here if needed

// // 	if err := c.CategoryRepository.CreateCategory(&category); err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Category"})
// // 		return
// // 	}
// // 	ctx.JSON(http.StatusCreated, category)
// // }

//v2
// package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryController struct {
// 	CategoryRepository *repositories.CategoryRepository
// }

// func NewCategoryController(categoryRepository *repositories.CategoryRepository) *CategoryController {
// 	return &CategoryController{
// 		CategoryRepository: categoryRepository,
// 	}
// }

// func (c *CategoryController) CreateFullCategory(ctx *gin.Context) {
// 	var req struct {
// 		ForeignID   uint   `json:"foreign_id"`
// 		Name        string `json:"name"`
// 		ParentID    uint   `json:"parent_id"`
// 		Description string `json:"description"`
// 		TenantID    string `json:"tenant_id"` // Assuming tenant_id is included in the JSON
// 	}

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Create models based on the request
// 	category := &models.Category{
// 		Tenant_ID: req.TenantID,
// 		Parent_ID: req.ParentID,
// 	}

// 	categoryLang := &models.CategoryLang{
// 		Tenant_ID:   req.TenantID,
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	categoryShop := &models.CategoryShop{
// 		Tenant_ID:  req.TenantID,
// 		Foreign_ID: req.ForeignID,
// 	}

// 	// Use the repository method to create the full category
// 	if err := c.CategoryRepository.CreateFullCategory(category, categoryLang, categoryShop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
// }

//v3 send data correct
// package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryController struct {
// 	CategoryRepository *repositories.CategoryRepository
// }

// func NewCategoryController(categoryRepository *repositories.CategoryRepository) *CategoryController {
// 	return &CategoryController{
// 		CategoryRepository: categoryRepository,
// 	}
// }

// func (c *CategoryController) CreateFullCategory(ctx *gin.Context) {
// 	var req struct {
// 		TenantID        string  `json:"tenant_id"`
// 		ParentID        uint    `json:"parent"`
// 		ForeignID       uint    `json:"foreign_id"`
// 		Name            string  `json:"name"`
// 		Description     string  `json:"description"`
// 		LevelDepth      *int    `json:"level_depth"`
// 		Active          *uint   `json:"active"`
// 		IsRootCategory  *uint   `json:"is_root_category"`
// 		Position        *int    `json:"position"`
// 		MetaTitle       *string `json:"meta_title"`
// 		MetaDescription *string `json:"meta_description"`
// 		MetaKeywords    *string `json:"meta_keywords"`
// 		LinkRewrite     *string `json:"link_rewrite"`
// 		LangID          uint    `json:"lang_id"` // New field for lang_id
// 		ShopID          uint    `json:"shop_id"` // New field for shop_id
// 	}

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Create models based on the request
// 	category := &models.Category{
// 		Tenant_ID: req.TenantID,
// 		Parent_ID: req.ParentID,
// 	}

// 	categoryLang := &models.CategoryLang{
// 		Tenant_ID:   req.TenantID,
// 		Lang_ID:     req.LangID,
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	categoryShop := &models.CategoryShop{
// 		Tenant_ID:  req.TenantID,
// 		Shop_ID:    req.ShopID, // Using the provided shop_id
// 		Foreign_ID: req.ForeignID,
// 	}

// 	// Use the repository method to create the full category
// 	if err := c.CategoryRepository.CreateFullCategory(category, categoryLang, categoryShop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
// }

//version working
// package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryController struct {
// 	CategoryRepository *repositories.CategoryRepository
// }

// func NewCategoryController(categoryRepository *repositories.CategoryRepository) *CategoryController {
// 	return &CategoryController{
// 		CategoryRepository: categoryRepository,
// 	}
// }

// func (c *CategoryController) CreateFullCategory(ctx *gin.Context) {
// 	var req struct {
// 		TenantID        string  `json:"tenant_id"`
// 		ParentID        uint    `json:"parent"`
// 		ForeignID       uint    `json:"foreign_id"`
// 		Name            string  `json:"name"`
// 		Description     string  `json:"description"`
// 		LevelDepth      *int    `json:"level_depth"`
// 		Active          *uint   `json:"active"`
// 		IsRootCategory  *uint   `json:"is_root_category"`
// 		Position        *int    `json:"position"`
// 		MetaTitle       *string `json:"meta_title"`
// 		MetaDescription *string `json:"meta_description"`
// 		MetaKeywords    *string `json:"meta_keywords"`
// 		LinkRewrite     *string `json:"link_rewrite"`
// 		LangID          uint    `json:"lang_id"` // New field for lang_id
// 		ShopID          uint    `json:"shop_id"` // New field for shop_id
// 	}

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Create models based on the request
// 	category := &models.Category{
// 		Tenant_ID: req.TenantID,
// 		Parent_ID: req.ParentID,
// 	}

// 	categoryLang := &models.CategoryLang{
// 		Tenant_ID:   req.TenantID,
// 		Lang_ID:     req.LangID,
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	categoryShop := &models.CategoryShop{
// 		Tenant_ID:  req.TenantID,
// 		Shop_ID:    req.ShopID, // Using the provided shop_id
// 		Foreign_ID: req.ForeignID,
// 	}

// 	// Use the repository method to create the full category
// 	if err := c.CategoryRepository.CreateFullCategory(category, categoryLang, categoryShop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
// }

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
