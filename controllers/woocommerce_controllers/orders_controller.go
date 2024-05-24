package woocommerce_controllers

import (
	"net/http"
	"shifti-connector-backend/models"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderRepository    *repositories.OrderRepository
	CustomerRepository *repositories.CustomerRepository
}

func NewOrderController(orderRepository *repositories.OrderRepository, customerRepository *repositories.CustomerRepository) *OrderController {
	return &OrderController{
		OrderRepository:    orderRepository,
		CustomerRepository: customerRepository,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Look up the actual customer_id based on foreign_id
	customerID, err := c.CustomerRepository.GetCustomerIDByForeignID(order.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find customer"})
		return
	}
	order.CustomerID = customerID

	if err := c.OrderRepository.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}
