package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {

	queries, conn, err := connectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close(context.Background())

	router := gin.Default()

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "AutoERP API is running",
		})
	})

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from AutoERP",
		})
	})

	router.POST("/api/customers", func(c *gin.Context) {
		createCustomer(c, queries)
	})

	router.GET("api/customers", func(c *gin.Context){
		getCustomers(c, queries)
	})

	router.GET("api/customers/:id", func(c *gin.Context){
		getCustomerByID(c,queries)
	})

	router.PUT("api/customers/:id", func(c *gin.Context){
		updateCustomerByID(c, queries)
	})

	router.DELETE("/api/customers/:id", func(c *gin.Context) {
		deleteCustomerByID(c, queries)
	})

	router.POST("/api/vehicles", func(c *gin.Context) {
		createVehicle(c, queries)
	})

	router.GET("/api/vehicles", func(c *gin.Context) {
		getVehicles(c, queries)
	})

	router.GET("/api/vehicles/:id", func(c *gin.Context) {
		getVehicleByID(c, queries)
	})

	router.GET("/api/customers/:id/vehicles", func(c *gin.Context) {
		getVehiclesByCustomerID(c, queries)
	})

	router.PUT("/api/vehicles/:id", func(c *gin.Context) {
		updateVehicleByID(c, queries)
	})
	
	router.DELETE("/api/vehicles/:id", func(c *gin.Context) {
		deleteVehicleByID(c, queries)
	})

	router.POST("/api/services", func(c *gin.Context) {
		createService(c, queries)
	})
	
	router.GET("/api/services", func(c *gin.Context) {
		getServices(c, queries)
	})
	
	router.GET("/api/services/:id", func(c *gin.Context) {
		getServiceByID(c, queries)
	})
	
	router.PUT("/api/services/:id", func(c *gin.Context) {
		updateServiceByID(c, queries)
	})
	
	router.DELETE("/api/services/:id", func(c *gin.Context) {
		deleteServiceByID(c, queries)
	})


	router.POST("/api/quotations", func(c *gin.Context) {
		createQuotation(c, queries)
	})

	router.GET("/api/quotations/:id", func(c *gin.Context) {
		getQuotationByID(c, queries)
	})

	router.GET("/api/quotations", func(c *gin.Context) {
		getQuotations(c, queries)
	})

	router.PUT("/api/quotations/:id/status", func(c *gin.Context) {
		updateQuotationStatus(c, queries)
	})

	router.DELETE("/api/quotations/:id", func(c *gin.Context) {
		deleteQuotationByID(c, queries)
	})


	router.POST("/api/work-orders", func(c *gin.Context) {
		createWorkOrder(c, queries)
	})
	
	router.GET("/api/work-orders", func(c *gin.Context) {
		getWorkOrders(c, queries)
	})
	
	router.GET("/api/work-orders/:id", func(c *gin.Context) {
		getWorkOrderByID(c, queries)
	})
	
	router.PUT("/api/work-orders/:id/status", func(c *gin.Context) {
		updateWorkOrderStatus(c, queries)
	})
	
	router.DELETE("/api/work-orders/:id", func(c *gin.Context) {
		deleteWorkOrderByID(c, queries)
	})

	router.POST("/api/invoices", func(c *gin.Context) {
		createInvoice(c, queries)
	})
	
	router.GET("/api/invoices", func(c *gin.Context) {
		getInvoices(c, queries)
	})

	router.GET("/api/invoices/:id", func(c *gin.Context) {
		getInvoiceByID(c, queries)
	})

	router.PUT("/api/invoices/:id/status", func(c *gin.Context) {
		updateInvoiceStatus(c, queries)
	})


    router.POST("/api/payments", func(c *gin.Context){
		createPayment(c, queries)
	})

	router.GET("/api/invoices/:id/payments", func(c *gin.Context) {
		getPaymentsByInvoice(c, queries)
	})

	router.GET("/api/payments", func(c *gin.Context) {
		getPayments(c, queries)
	})

	router.GET("/api/payments/:id", func(c *gin.Context) {
		getPaymentByID(c, queries)
	})

	router.DELETE("/api/payments/:id", func(c *gin.Context) {
		deletePaymentByID(c, queries)
	})

	router.GET("/api/work-orders/status/:status", func(c *gin.Context) {
		getWorkOrdersByStatus(c, queries)
	})

	router.GET("/api/quotations/status/:status", func(c *gin.Context) {
		getQuotationsByStatus(c, queries)
	})

	router.GET("/api/invoices/status/:status", func(c *gin.Context) {
		getInvoicesByStatus(c, queries)
	})

	router.GET("/api/dashboard", func(c *gin.Context) {
		getDashboard(c, queries)
	})

	router.Run(":8080")

	_ = queries
}