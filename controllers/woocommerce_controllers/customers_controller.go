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
		// Address fields
		ManufacturerID uint   `json:"manufacturer_id"`
		SupplierID     uint   `json:"supplier_id"`
		WarehouseID    uint   `json:"warehouse_id"`
		CountryID      uint   `json:"country_id"`
		StateID        uint   `json:"state_id"`
		Alias          string `json:"alias"`
		Company        string `json:"company"`
		VatNumber      string `json:"vat_number"`
		Address1       string `json:"address1"`
		Address2       string `json:"address2"`
		Postcode       string `json:"postcode"`
		City           string `json:"city"`
		Other          string `json:"other"`
		Phone          string `json:"phone"`
		PhoneMobile    string `json:"phone_mobile"`
		DNI            string `json:"dni"`
		WooBilling     string `json:"woo_billing"`
		WooShipping    string `json:"woo_shipping"`
		WooCountry     string `json:"woo_country"`
		WooState       string `json:"woo_state"`
	}

	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		customer := &models.Customer{
			Tenant_ID:  req.TenantID,
			Foreign_ID: req.ForeignID,
			Shop_ID:    req.ShopID,
			First_Name: req.FirstName,
			Last_Name:  req.LastName,
			Email:      req.Email,
		}

		customerShop := &models.CustomerShop{
			Tenant_ID:  req.TenantID,
			Shop_ID:    req.ShopID,
			Foreign_ID: req.ForeignID,
		}

		address := &models.Addresses{
			Tenant_ID:       req.TenantID,
			Shop_ID:         req.ShopID,
			Foreign_ID:      req.ForeignID,
			Manufacturer_ID: req.ManufacturerID,
			Supplier_ID:     req.SupplierID,
			Warehouse_ID:    req.WarehouseID,
			Country_ID:      req.CountryID,
			State_ID:        req.StateID,
			Alias:           req.Alias,
			Company:         req.Company,
			Last_Name:       req.LastName,
			First_Name:      req.FirstName,
			VAT_Number:      req.VatNumber,
			Address1:        req.Address1,
			Address2:        req.Address2,
			Postcode:        req.Postcode,
			City:            req.City,
			Other:           req.Other,
			Phone:           req.Phone,
			Phone_Mobile:    req.PhoneMobile,
			DNI:             req.DNI,

			Woo_Country: req.WooCountry,
			Woo_State:   req.WooState,
		}

		if err := c.CustomerRepository.CreateFullCustomer(customer, customerShop, address); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer"})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Customers created successfully"})
}
