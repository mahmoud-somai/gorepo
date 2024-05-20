package main

import (
	"fmt"
	"os"
	"shifti-connector-backend/config"
	"shifti-connector-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	// Configure CORS
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"*"} // Permettre à tous les origines
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(configCors))

	// configure trusted proxies
	// r.SetTrustedProxies([]string{"127.0.0.1"}) // Trusted proxies configuration

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Perform automatic migration using the migration package
	// if err := migration.MigrateDB(db); err != nil {
	//     log.Fatalf("Failed to migrate database: %v", err)
	// }

	// Initialize Gin router
	routes.SetupPrestashopRoutes(r, db)
	routes.SetupWoocommerceRoutes(r, db)

	// Start the server
	port := os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
