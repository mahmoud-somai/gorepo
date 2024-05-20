// func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
// 	var customer models.Customer

// 	if err := ctx.ShouldBindJSON(&customer); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CustomerRepository.CreateCustomer(&customer); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, customer)
// }

// package woocommerce_controllers

// import (
// 	"net/http"
// 	"shifti-connector-backend/models"
// 	"shifti-connector-backend/repositories"

// 	"github.com/gin-gonic/gin"
// )

// type CustomerController struct {
// 	CustomerRepository *repositories.CustomerRepository
// }

// func NewCustomerRepository(customerRepository *repositories.CustomerRepository) *CustomerController {
// 	return &CustomerController{
// 		CustomerRepository: customerRepository,
// 	}
// }

// func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
// 	var customers []models.Customer

// 	if err := ctx.ShouldBindJSON(&customers); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	for _, customer := range customers {
// 		if err := c.CustomerRepository.CreateCustomer(&customer); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer"})
// 			return
// 		}
// 	}

// 	ctx.JSON(http.StatusCreated, customers)
// }

// func (c *CustomerController) CreateCustomerMessage(ctx *gin.Context) {
// 	var customer_message models.CustomerMessage

// 	if err := ctx.ShouldBindJSON(&customer_message); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CustomerRepository.CreateCustomerMessage(&customer_message); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer Message"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, customer_message)
// }

// func (c *CustomerController) CreateCustomerShop(ctx *gin.Context) {
// 	var customer_shop models.CustomerShop

// 	if err := ctx.ShouldBindJSON(&customer_shop); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CustomerRepository.CreateCustomerShop(&customer_shop); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer Shop"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, customer_shop)
// }

// func (c *CustomerController) CreateCustomerThread(ctx *gin.Context) {
// 	var customer_thread models.CustomerThread

// 	if err := ctx.ShouldBindJSON(&customer_thread); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// You might want to validate the example data here if needed

// 	if err := c.CustomerRepository.CreateCustomerThread(&customer_thread); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer Thread"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, customer_thread)
// }

package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerRepository *repositories.CustomerRepository
}

func NewCustomerController(customerRepository *repositories.CustomerRepository) *CustomerController {
	return &CustomerController{
		CustomerRepository: customerRepository,
	}
}

func (c *CustomerController) CreateFullCustomers(ctx *gin.Context) {
	var reqs []struct {
		TenantID  string `json:"tenant_id"`
		ForeignID uint   `json:"foreign_id"`
		ShopID    uint   `json:"shop_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdCustomers []models.Customer

	for _, req := range reqs {
		// Create the customer model
		customer := &models.Customer{
			Tenant_ID:  req.TenantID,
			Foreign_ID: req.ForeignID,
			Shop_ID:    req.ShopID,
			First_Name: req.FirstName,
			Last_Name:  req.LastName,
			Email:      req.Email,
		}

		// Create the customer shop model
		customerShop := &models.CustomerShop{
			Tenant_ID:  req.TenantID,
			Shop_ID:    req.ShopID,
			Foreign_ID: req.ForeignID, // Same foreign_id as in the customer
		}

		// Use the repository method to create the full customer
		if err := c.CustomerRepository.CreateFullCustomer(customer, customerShop); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer"})
			return
		}

		createdCustomers = append(createdCustomers, *customer)
	}

	ctx.JSON(http.StatusCreated, createdCustomers)
}
