package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"shifti-connector-backend/repositories"
    "shifti-connector-backend/controllers/prestashop_controllers"
)

// SetupRoutes initializes the API routes and associates them with handler functions.
func SetupPrestashopRoutes(r *gin.Engine, db *gorm.DB) {
    exampleRepository := repositories.NewExampleRepository(db)
    exampleController := prestashop_controllers.NewExampleController(exampleRepository)

    // Example routes
    exampleGroup := r.Group("/prestashop/example")
    {
        exampleGroup.POST("/", exampleController.CreateExample)
        exampleGroup.GET("/:id", exampleController.GetExampleByID)
        exampleGroup.PUT("/:id", exampleController.UpdateExample)
        exampleGroup.DELETE("/:id", exampleController.DeleteExampleByID)
    }

    // Add more routes as needed
}
