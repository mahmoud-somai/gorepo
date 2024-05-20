package prestashop_controllers

import (
    "github.com/gin-gonic/gin"
    "shifti-connector-backend/models"
	"shifti-connector-backend/DTOs/prestashop_DTOs"
    "shifti-connector-backend/repositories"
    "net/http"
    "strconv"
)

// ExampleController represents the controller for Example entity.
type ExampleController struct {
    ExampleRepository *repositories.ExampleRepository
}

// NewExampleController creates a new instance of ExampleController.
func NewExampleController(exampleRepository *repositories.ExampleRepository) *ExampleController {
    return &ExampleController{
        ExampleRepository: exampleRepository,
    }
}

// CreateExample handles creating a new example entity.
func (c *ExampleController) CreateExample(ctx *gin.Context) {

	var exampleDto prestashop_DTOs.ExampleDTO
    var example models.Example

    if err := ctx.ShouldBindJSON(&exampleDto); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// handle data
    example=*exampleDto.ToModel()
	
    if err := c.ExampleRepository.CreateExample(&example); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create example"})
        return
    }
    ctx.JSON(http.StatusCreated, example)
}

// GetExampleByID handles retrieving an example entity by ID.
func (c *ExampleController) GetExampleByID(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    example, err := c.ExampleRepository.GetExampleByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Example not found"})
        return
    }
    ctx.JSON(http.StatusOK, example)
}

// UpdateExample handles updating an example entity.
func (c *ExampleController) UpdateExample(ctx *gin.Context) {
    var example models.Example
    if err := ctx.ShouldBindJSON(&example); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.ExampleRepository.UpdateExample(&example); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update example"})
        return
    }
    ctx.JSON(http.StatusOK, example)
}

// DeleteExampleByID handles deleting an example entity by ID.
func (c *ExampleController) DeleteExampleByID(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.ExampleRepository.DeleteExampleByID(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete example"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Example deleted successfully"})
}
