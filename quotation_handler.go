package main

import (
	"net/http"
	"strconv"

	db "autoerp/db/sqlc"

	"github.com/gin-gonic/gin"
)
func createQuotation(c *gin.Context, queries *db.Queries) {
	var request CreateQuotationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quotation data",
		})
		return
	}

	if len(request.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quotation must contain at least one item",
		})
		return
	}

	var total float64

	// First calculate total using service prices
	for _, item := range request.Items {

		service, err := queries.GetService(c, int32(item.ServiceID))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Service not found",
			})
			return
		}

		amount := service.Price * float64(item.Quantity)
		total += amount
	}

	// Create quotation
	quotation, err := queries.CreateQuotation(
		c,
		db.CreateQuotationParams{
			CustomerID:  int32(request.CustomerID),
			VehicleID:   int32(request.VehicleID),
			Status:    QuotationDraft,
			TotalAmount: total,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create quotation",
		})
		return
	}

	// Add quotation items
	for _, item := range request.Items {

		service, err := queries.GetService(c, int32(item.ServiceID))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Service not found",
			})
			return
		}

		amount := service.Price * float64(item.Quantity)

		_, err = queries.AddQuotationItem(
			c,
			db.AddQuotationItemParams{
				QuotationID: int32(quotation.ID),
				ServiceID:   int32(item.ServiceID),
				Quantity:    int32(item.Quantity),
				Price:       service.Price,
				Amount:      amount,
			},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add quotation item",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"quotation": quotation,
		"message":   "Quotation created successfully",
	})
}


func getQuotationByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quotation ID",
		})
		return
	}

	quotation, err := queries.GetQuotation(c, int32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Quotation not found",
		})
		return
	}

	items, err := queries.GetQuotationItems(c, int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get quotation items",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quotation": quotation,
		"items":     items,
	})
}

func getQuotations(c *gin.Context, queries *db.Queries) {
    quotations, err := queries.ListQuotations(c)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get quotations",
        })
        return
    }

    c.JSON(http.StatusOK, quotations)
}


func updateQuotationStatus(c *gin.Context, queries *db.Queries) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid quotation ID",
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

    quotation, err := queries.UpdateQuotationStatus(
        c,
        db.UpdateQuotationStatusParams{
            Status: request.Status,
            ID:     int32(id),
        },
    )

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Quotation not found",
        })
        return
    }

    c.JSON(http.StatusOK, quotation)
}

func deleteQuotationByID(c *gin.Context, queries *db.Queries) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid quotation ID",
        })
        return
    }

    err = queries.DeleteQuotation(c, int32(id))

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to delete quotation",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Quotation deleted successfully",
    })
}


func getQuotationsByStatus(c *gin.Context, queries *db.Queries) {
	status := c.Param("status")

	if status != QuotationDraft &&
		status != QuotationApproved &&
		status != QuotationRejected {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quotation status"})
		return
	}

	quotations, err := queries.GetQuotationsByStatus(c, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get quotations"})
		return
	}

	if quotations == nil {
		quotations = []db.Quotation{}
	}

	c.JSON(http.StatusOK, quotations)
}