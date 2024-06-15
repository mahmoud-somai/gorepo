package routes

import (
	"shifti-connector-backend/controllers/woocommerce_controllers"
	"shifti-connector-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes initializes the API routes and associates them with handler functions.
func SetupWoocommerceRoutes(r *gin.Engine, db *gorm.DB) {
	// exampleRepository := repositories.NewExampleRepository(db)
	// exampleController := woocommerce_controllers.NewExampleController(exampleRepository)

	// Example routes
	// exampleGroup := r.Group("/woocommerce/example")
	// {
	//     exampleGroup.POST("/", exampleController.CreateExample)
	//     exampleGroup.GET("/:id", exampleController.GetExampleByID)
	//     exampleGroup.PUT("/:id", exampleController.UpdateExample)
	//     exampleGroup.DELETE("/:id", exampleController.DeleteExampleByID)
	// }

	// Add more routes as needed
	// billingRepository := repositories.NewBillingRepository(db)
	// billingController := woocommerce_controllers.BillingController(billingRepository)

	attributeRepository := repositories.NewAttributeRepository(db)
	attributeController := woocommerce_controllers.NewAttributeRepository(attributeRepository)

	addressRepository := repositories.NewAddressesRepository(db)
	addressController := woocommerce_controllers.NewAddressRepository(addressRepository)

	attributeGroupRepository := repositories.NewAttributeGroupRepository(db)
	attributeGroupController := woocommerce_controllers.NewAttributeGroupRepository(attributeGroupRepository)

	attributeGroupLangRepository := repositories.NewAttributeGroupLangRepository(db)
	attributeGroupLangController := woocommerce_controllers.NewAttributeGroupLangRepository(attributeGroupLangRepository)

	attributeGroupShopRepository := repositories.NewAttributeGroupShopRepository(db)
	attributeGroupShopController := woocommerce_controllers.NewAttributeGroupShopRepository(attributeGroupShopRepository)

	attributeLangRepository := repositories.NewAttributeLangRepository(db)
	attributeLangController := woocommerce_controllers.NewAttributeLangRepository(attributeLangRepository)

	attributeShopRepository := repositories.NewAttributeShopRepository(db)
	attributeShopController := woocommerce_controllers.NewAttributeShopRepository(attributeShopRepository)

	carrierRepository := repositories.NewCarrierRepository(db)
	carrierController := woocommerce_controllers.NewCarrierRepository(carrierRepository)

	carrierLangRepository := repositories.NewCarrierLangRepository(db)
	carrierLangController := woocommerce_controllers.NewCarrierLangRepository(carrierLangRepository)

	carrierShopRepository := repositories.NewCarrierShopRepository(db)
	carrierShopController := woocommerce_controllers.NewCarrierShopRepository(carrierShopRepository)

	carrierZoneRepository := repositories.NewCarrierZoneRepository(db)
	carrierZoneController := woocommerce_controllers.NewCarrierZoneRepository(carrierZoneRepository)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryController := woocommerce_controllers.NewCategoryController(categoryRepository)

	// categoryLangRepository := repositories.NewCategoryLangRepository(db)
	// categoryLangController := woocommerce_controllers.NewCategoryLangRepository(categoryLangRepository)

	// categoryShopRepository := repositories.NewCategoryShopRepository(db)
	// categoryShopController := woocommerce_controllers.NewCategoryShopRepository(categoryShopRepository)

	// categoryProductRepository := repositories.NewCategoryProductRepository(db)
	// categoryProductController := woocommerce_controllers.NewCategoryProductRepository(categoryProductRepository)

	countryRepository := repositories.NewCountryRepository(db)
	countryController := woocommerce_controllers.NewCountryRepository(countryRepository)

	// customerRepository := repositories.NewCustomerRepository(db)
	// customerController := woocommerce_controllers.NewCustomerRepository(customerRepository)

	customerRepository := repositories.NewCustomerRepository(db)
	customerController := woocommerce_controllers.NewCustomerController(customerRepository)

	taxRepository := repositories.NewTaxeRepository(db)
	TaxController := woocommerce_controllers.NewTaxController(taxRepository)

	productRepository := repositories.NewProductRepository(db)
	ProductController := woocommerce_controllers.NewProductController(productRepository)

	deliveryRepository := repositories.NewDeliveryRepository(db)
	deliveryController := woocommerce_controllers.NewDeliveryRepository(deliveryRepository)

	featureRepository := repositories.NewFeatureRepository(db)
	featureController := woocommerce_controllers.NewFeatureRepository(featureRepository)

	orderRepository := repositories.NewOrderRepository(db)
	orderController := woocommerce_controllers.NewOrderController(orderRepository, customerRepository)

	orderDetailsRepository := repositories.NewOrderDetailsRepository(db)
	orderDetailsController := woocommerce_controllers.NewOrderDetailController(orderDetailsRepository)

	orderFeesRepository := repositories.NewOrderFeeRepository(db)
	orderFeesController := woocommerce_controllers.NewOrderFeeController(orderFeesRepository)

	orderCarriersRepository := repositories.NewOrderCarriersRepository(db)
	orderCarriersController := woocommerce_controllers.NewOrderCarrierController(orderCarriersRepository)

	orderTaxesRepository := repositories.NewOrderTaxRepository(db)
	orderTaxesController := woocommerce_controllers.NewOrderTaxController(orderTaxesRepository)

	orderCarrierTaxRespository := repositories.NewOrderCarrierTaxRepository(db)
	orderCarriersTaxController := woocommerce_controllers.NewOrderCarrierTaxController(orderCarrierTaxRespository)

	orderDetailTaxRepository := repositories.NewOrderDetailsTaxRepository(db)
	orderDetailsTaxController := woocommerce_controllers.NewOrderDetailsTaxController(orderDetailTaxRepository)

	billingRepository := repositories.NewBillingRepository(db)
	billingController := woocommerce_controllers.NewBillingController(billingRepository)

	shippingRepository := repositories.NewShippingRepository(db)
	shippingController := woocommerce_controllers.NewShippingController(shippingRepository)

	shopRepo := repositories.NewShopRepository(db)
	shopController := woocommerce_controllers.NewShopController(shopRepo)

	woocommerceGroup := r.Group("/woocommerce")
	{
		woocommerceGroup.POST("/billing", billingController.CreateBilling)
		woocommerceGroup.POST("/shipping", shippingController.CreateShipping)

		woocommerceGroup.POST("/address", addressController.CreateAddress)

		woocommerceGroup.POST("/attribute", attributeController.CreateAttribute)
		woocommerceGroup.POST("/attributeGroup", attributeGroupController.CreateAttributeGroup)
		woocommerceGroup.POST("/attributeGroupLang", attributeGroupLangController.CreateAttributeGroupLang)
		woocommerceGroup.POST("/attributeGroupShop", attributeGroupShopController.CreateAttributeGroupShop)
		woocommerceGroup.POST("/attributeLang", attributeLangController.CreateAttributeLang)
		woocommerceGroup.POST("/attributeShop", attributeShopController.CreateAttributeShop)

		woocommerceGroup.POST("/carrier", carrierController.CreateCarrier)
		woocommerceGroup.POST("/carrierLang", carrierLangController.CreateCarrierLang)
		woocommerceGroup.POST("/carrierShop", carrierShopController.CreateCarrierShop)
		woocommerceGroup.POST("/carrierZone", carrierZoneController.CreateCarrierZone)

		woocommerceGroup.POST("/category", categoryController.CreateFullCategory)
		// woocommerceGroup.POST("/categoryProduct", categoryController.CreateCategoryProduct)
		// woocommerceGroup.POST("/categoryLang", categoryController.CreateCategoryLang)
		// woocommerceGroup.POST("/categoryShop", categoryController.CreateCategoryShop)

		woocommerceGroup.POST("/country", countryController.CreateCountry)
		woocommerceGroup.POST("/countryLang", countryController.CreateCountryLang)
		woocommerceGroup.POST("/countryShop", countryController.CreateCountryShop)

		// woocommerceGroup.POST("/customer", customerController.CreateCustomer)
		// woocommerceGroup.POST("/customerMessage", customerController.CreateCustomerMessage)
		// woocommerceGroup.POST("/customerShop", customerController.CreateCustomerShop)
		// woocommerceGroup.POST("/customerThread", customerController.CreateCustomerThread)

		woocommerceGroup.POST("/customer", customerController.CreateFullCustomers)
		woocommerceGroup.POST("/taxe", TaxController.CreateTaxes)
		woocommerceGroup.POST("/product", ProductController.CreateFullProducts)

		woocommerceGroup.POST("/delivery", deliveryController.CreateDelivery)

		woocommerceGroup.POST("/feature", featureController.CreateFeature)
		woocommerceGroup.POST("/featureLang", featureController.CreateFeatureLang)
		woocommerceGroup.POST("/featureShop", featureController.CreateFeatureShop)
		woocommerceGroup.POST("/featureValue", featureController.CreateFeatureValue)
		woocommerceGroup.POST("/featureValueLang", featureController.CreateFeatureValueLang)

		woocommerceGroup.POST("/order", orderController.CreateOrders)
		woocommerceGroup.POST("/orderdetails", orderDetailsController.CreateOrderDetails)
		woocommerceGroup.POST("/orderFees", orderFeesController.CreateOrderFees)
		woocommerceGroup.POST("/orderCarriers", orderCarriersController.CreateOrderCarrier)
		woocommerceGroup.POST("/orderTaxes", orderTaxesController.CreateOrderTax)
		woocommerceGroup.POST("/orderCarrierTax", orderCarriersTaxController.CreateOrderCarrierTaxes)
		woocommerceGroup.POST("/orderDetailsTax", orderDetailsTaxController.CreateOrderDetailsTaxes)
		woocommerceGroup.GET("/shop", shopController.GetShopByURL)

	}
}
