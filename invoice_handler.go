package main

import (
	"net/http"
	"strconv"

	db "autoerp/db/sqlc"

	"github.com/gin-gonic/gin"
)

/*And we're using SQLC functions that already exist:

queries.GetWorkOrder(...)
queries.GetQuotation(...)
queries.CreateInvoice(...)*/

func createInvoice(c *gin.Context, queries *db.Queries) {
	var request CreateInvoiceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice data",
		})
		return
	}

	// Get work order
	workOrder, err := queries.GetWorkOrder(
		c,
		int32(request.WorkOrderID),
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Work order not found",
		})
		return
	}

	// Get quotation
	quotation, err := queries.GetQuotation(
		c,
		workOrder.QuotationID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Quotation not found",
		})
		return
	}

	// Create invoice
	invoice, err := queries.CreateInvoice(
		c,
		db.CreateInvoiceParams{
			WorkOrderID: int32(request.WorkOrderID),
			CustomerID:  quotation.CustomerID,
			TotalAmount: quotation.TotalAmount,
			Status:      InvoiceUnpaid,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create invoice",
		})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func getInvoices(c *gin.Context, queries *db.Queries) {
	invoices, err := queries.ListInvoices(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get invoices",
		})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

func getInvoiceByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice ID",
		})
		return
	}

	invoice, err := queries.GetInvoice(c, int32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invoice not found",
		})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func updateInvoiceStatus(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice ID",
		})
		return
	}

	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status",
		})
		return
	}

	// Validate invoice status
	if request.Status != InvoiceUnpaid &&
		request.Status != InvoicePartial &&
		request.Status != InvoicePaid {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice status",
		})
		return
	}

	invoice, err := queries.UpdateInvoiceStatus(
		c,
		db.UpdateInvoiceStatusParams{
			Status: request.Status,
			ID:     int32(id),
		},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invoice not found",
		})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func getInvoicesByStatus(c *gin.Context, queries *db.Queries) {
	status := c.Param("status")

	if status != InvoiceUnpaid &&
		status != InvoicePartial &&
		status != InvoicePaid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice status"})
		return
	}

	invoices, err := queries.GetInvoicesByStatus(c, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get invoices"})
		return
	}

	if invoices == nil {
		invoices = []db.Invoice{}
	}

	c.JSON(http.StatusOK, invoices)
}