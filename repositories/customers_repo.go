package repositories

import (
	"shifti-connector-backend/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		DB: db,
	}
}

func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	return r.DB.Create(customer).Error
}

func (r *CustomerRepository) CreateCustomerShop(customerShop *models.CustomerShop) error {
	return r.DB.Create(customerShop).Error
}

func (r *CustomerRepository) CreateAddress(address *models.Addresses) error {
	return r.DB.Create(address).Error
}

func (r *CustomerRepository) CreateFullCustomer(customer *models.Customer, customerShop *models.CustomerShop, address *models.Addresses) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert into customers table
	if err := tx.Create(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the CustomerID for customerShop and address
	customerShop.Customer_ID = customer.ID
	address.Customer_ID = customer.ID

	// Insert into customer_shops table
	if err := tx.Create(customerShop).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert into addresses table
	if err := tx.Create(address).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *CustomerRepository) GetCustomerIDByForeignID(foreignID int32) (int32, error) {
	var customer models.Customer
	if err := r.DB.Where("foreign_id = ?", foreignID).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, err
		}
		return 0, err
	}
	return int32(customer.ID), nil
}
