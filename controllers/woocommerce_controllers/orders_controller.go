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

func (c *OrderController) CreateOrders(ctx *gin.Context) {
	var orders []models.Order

	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var failedOrders []models.Order
	var failedReasons []string

	for _, order := range orders {
		// Look up the actual customer_id based on foreign_id
		customerID, err := c.CustomerRepository.GetCustomerIDByForeignID(order.CustomerID)
		if err != nil {
			failedOrders = append(failedOrders, order)
			failedReasons = append(failedReasons, "Failed to find customer")
			continue
		}
		order.CustomerID = customerID

		if err := c.OrderRepository.CreateOrder(&order); err != nil {
			failedOrders = append(failedOrders, order)
			failedReasons = append(failedReasons, "Failed to create order")
			continue
		}
	}

	if len(failedOrders) > 0 {
		ctx.JSON(http.StatusPartialContent, gin.H{
			"message":       "Some orders failed to process",
			"failed_orders": failedOrders,
			"reasons":       failedReasons,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "All orders created successfully"})
}
