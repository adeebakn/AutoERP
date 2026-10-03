package main

import (
	"net/http"
	"strconv"
	"time"

	db "autoerp/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)
func createWorkOrder(c *gin.Context, queries *db.Queries) {
	var request CreateWorkOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order data",
		})
		return
	}

	dueDate, err := time.Parse("2006-01-02", request.DueDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid due date. Use YYYY-MM-DD",
		})
		return
	}

	workOrder, err := queries.CreateWorkOrder(
		c,
		db.CreateWorkOrderParams{
			QuotationID: int32(request.QuotationID),
			CustomerID:  int32(request.CustomerID),
			VehicleID:   int32(request.VehicleID),
			Status:      WorkOrderPending,
			DueDate: pgtype.Date{
				Time:  dueDate,
				Valid: true,
			},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create work order",
		})
		return
	}

	c.JSON(http.StatusCreated, workOrder)
}

func getWorkOrders(c *gin.Context, queries *db.Queries) {
	workOrders, err := queries.ListWorkOrders(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get work orders",
		})
		return
	}

	c.JSON(http.StatusOK, workOrders)
}

func getWorkOrderByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order ID",
		})
		return
	}

	workOrder, err := queries.GetWorkOrder(c, int32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Work order not found",
		})
		return
	}

	c.JSON(http.StatusOK, workOrder)
}

func updateWorkOrderStatus(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order ID",
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
	
	// Validate status
	if request.Status != WorkOrderPending &&
		request.Status != WorkOrderInProgress &&
		request.Status != WorkOrderCompleted &&
		request.Status != WorkOrderCancelled {
	
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order status",
		})
		return
	}
	
	// Update database
	workOrder, err := queries.UpdateWorkOrderStatus(
		c,
		db.UpdateWorkOrderStatusParams{
			Status: request.Status,
			ID:     int32(id),
		},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Work order not found",
		})
		return
	}

	c.JSON(http.StatusOK, workOrder)
}

func deleteWorkOrderByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order ID",
		})
		return
	}

	err = queries.DeleteWorkOrder(c, int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete work order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Work order deleted successfully",
	})
}


func getWorkOrdersByStatus(c *gin.Context, queries *db.Queries) {
	status := c.Param("status")

	if status != WorkOrderPending &&
		status != WorkOrderInProgress &&
		status != WorkOrderCompleted &&
		status != WorkOrderCancelled {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid work order status",
		})
		return
	}

	workOrders, err := queries.GetWorkOrdersByStatus(c, status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get work orders",
		})
		return
	}

	if workOrders == nil {
		workOrders = []db.WorkOrder{}
	}

	c.JSON(http.StatusOK, workOrders)
}