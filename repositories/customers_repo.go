// package repositories

// import (
// 	"shifti-connector-backend/models"

// 	"gorm.io/gorm"
// )

// type CustomerRepository struct {
// 	DB *gorm.DB
// }

// func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
// 	return &CustomerRepository{
// 		DB: db,
// 	}
// }
// func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
// 	return r.DB.Create(customer).Error
// }

// func (r *CustomerRepository) CreateCustomerMessage(customer_message *models.CustomerMessage) error {
// 	return r.DB.Create(customer_message).Error
// }

// func (r *CustomerRepository) CreateCustomerShop(customer_shop *models.CustomerShop) error {
// 	return r.DB.Create(customer_shop).Error
// }

// func (r *CustomerRepository) CreateCustomerThread(customer_thread *models.CustomerThread) error {
// 	return r.DB.Create(customer_thread).Error
// }

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

func (r *CustomerRepository) CreateCustomerMessage(customer_message *models.CustomerMessage) error {
	return r.DB.Create(customer_message).Error
}

func (r *CustomerRepository) CreateCustomerShop(customer_shop *models.CustomerShop) error {
	return r.DB.Create(customer_shop).Error
}

func (r *CustomerRepository) CreateCustomerThread(customer_thread *models.CustomerThread) error {
	return r.DB.Create(customer_thread).Error
}

// CreateFullCustomer handles the creation of a customer and customer shop within a transaction
func (r *CustomerRepository) CreateFullCustomer(customer *models.Customer, customerShop *models.CustomerShop) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert into customers table
	if err := tx.Create(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the CustomerID for customerShop
	customerShop.Customer_ID = customer.ID

	// Insert into customer_shops table
	if err := tx.Create(customerShop).Error; err != nil {
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
