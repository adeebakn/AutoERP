package main

import (
	"net/http"
	"strconv"

	db "autoerp/db/sqlc"

	"github.com/gin-gonic/gin"
)

func createPayment(c *gin.Context, queries *db.Queries) {
	var request CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment data",
		})
		return
	}

	// Get invoice
	invoice, err := queries.GetInvoice(
		c,
		int32(request.InvoiceID),
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invoice not found",
		})
		return
	}

	// Payment must be greater than 0
	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment amount must be greater than 0",
		})
		return
	}

	if request.PaymentMethod != PaymentCash &&
		request.PaymentMethod != PaymentUPI &&
		request.PaymentMethod != PaymentCard &&
		request.PaymentMethod != PaymentBankTransfer {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment method",
		})
		return
	}

	// Get previous payments for this invoice
	payments, err := queries.GetPaymentsByInvoice(
		c,
		int32(request.InvoiceID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get previous payments",
		})
		return
	}

	// Calculate total already paid
	var totalPaid float64

	for _, payment := range payments {
		totalPaid += payment.Amount
	}

	// Calculate remaining balance
	remainingBalance := invoice.TotalAmount - totalPaid

	// Check if payment is greater than remaining balance
	if request.Amount > remainingBalance {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment amount exceeds remaining balance",
			"remaining_balance": remainingBalance,
		})
		return
	}

	// Create payment
	payment, err := queries.CreatePayment(
		c,
		db.CreatePaymentParams{
			InvoiceID:     int32(request.InvoiceID),
			Amount:        request.Amount,
			PaymentMethod: request.PaymentMethod,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create payment",
		})
		return
	}

	// Calculate balance after this payment
	remainingBalance -= request.Amount

	// Determine invoice status automatically
	invoiceStatus := InvoicePartial

	if remainingBalance == 0 {
		invoiceStatus = InvoicePaid
	}

	// Update invoice status
	updatedInvoice, err := queries.UpdateInvoiceStatus(
		c,
		db.UpdateInvoiceStatusParams{
			Status: invoiceStatus,
			ID:     invoice.ID,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Payment created but failed to update invoice status",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"payment":           payment,
		"invoice":           updatedInvoice,
		"invoice_total":     invoice.TotalAmount,
		"total_paid":        totalPaid + request.Amount,
		"remaining_balance": remainingBalance,
	})
}


func getPaymentsByInvoice(c *gin.Context, queries *db.Queries) {
	invoiceID, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice ID",
		})
		return
	}

	payments, err := queries.GetPaymentsByInvoice(
		c,
		int32(invoiceID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get payment history",
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func getPayments(c *gin.Context, queries *db.Queries) {
	payments, err := queries.ListPayments(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get payments",
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func getPaymentByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment ID",
		})
		return
	}

	payment, err := queries.GetPayment(c, int32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment not found",
		})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func deletePaymentByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment ID",
		})
		return
	}

	// Get payment before deleting it
	payment, err := queries.GetPayment(c, int32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment not found",
		})
		return
	}

	// Delete payment
	err = queries.DeletePayment(c, int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete payment",
		})
		return
	}

	// Get invoice
	invoice, err := queries.GetInvoice(c, payment.InvoiceID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invoice not found",
		})
		return
	}

	// Get remaining payments
	payments, err := queries.GetPaymentsByInvoice(
		c,
		payment.InvoiceID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get payment history",
		})
		return
	}

	// Calculate total paid
	var totalPaid float64

	for _, p := range payments {
		totalPaid += p.Amount
	}

	// Determine invoice status
	invoiceStatus := InvoiceUnpaid

	if totalPaid > 0 && totalPaid < invoice.TotalAmount {
		invoiceStatus = InvoicePartial
	}

	if totalPaid == invoice.TotalAmount {
		invoiceStatus = InvoicePaid
	}

	// Update invoice status
	_, err = queries.UpdateInvoiceStatus(
		c,
		db.UpdateInvoiceStatusParams{
			Status: invoiceStatus,
			ID:     invoice.ID,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Payment deleted but failed to update invoice status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Payment deleted successfully",
		"total_paid": totalPaid,
		"remaining_balance": invoice.TotalAmount - totalPaid,
		"invoice_status": invoiceStatus,
	})
}