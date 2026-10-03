package main

import(
	"net/http"
	"strconv"
	"fmt"
	db "autoerp/db/sqlc"
	"github.com/gin-gonic/gin"
	//"github.com/jackc/pgx/v5/pgtype"
)


func createVehicle(c *gin.Context, queries *db.Queries) {

	var vehicle Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Vehicle not found",
		})
		return
	}

	fmt.Printf("Vehicle received: %+v\n", vehicle)

	createdVehicle, err := queries.CreateVehicle(
		c,
		db.CreateVehicleParams{
			CustomerID:  int32(vehicle.CustomerID),
			RegNo:       vehicle.RegNo,
			VehicleType: vehicle.VehicleType,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdVehicle)
}


//lsit all vehicles

func getVehicles(c *gin.Context, queries *db.Queries) {

    vehicles, err := queries.ListVehicles(c)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get vehicles",
        })
        return
    }

    c.JSON(http.StatusOK, vehicles)
}

//get  vehicle by id
func getVehicleByID(c *gin.Context, queries *db.Queries) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid vehicle ID",
        })
        return
    }

    vehicle, err := queries.GetVehicle(c, int32(id))

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Vehicle not found",
        })
        return
    }

    c.JSON(http.StatusOK, vehicle)
}


//get the vehicle by customer id

func getVehiclesByCustomerID(c *gin.Context, queries *db.Queries) {
    customerID, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid customer ID",
        })
        return
    }

    vehicles, err := queries.GetVehiclesByCustomerID(c, int32(customerID))

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get vehicles",
        })
        return
    }

    c.JSON(http.StatusOK, vehicles)
}


// update vehicle 

func updateVehicleByID(c *gin.Context, queries *db.Queries) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid vehicle ID",
        })
        return
    }

    var vehicle Vehicle

    if err := c.ShouldBindJSON(&vehicle); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid vehicle data",
        })
        return
    }

    updatedVehicle, err := queries.UpdateVehicle(
        c,
        db.UpdateVehicleParams{
            CustomerID:  int32(vehicle.CustomerID),
            RegNo:       vehicle.RegNo,
            VehicleType: vehicle.VehicleType,
            ID:          int32(id),
        },
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to update vehicle",
        })
        return
    }

    c.JSON(http.StatusOK, updatedVehicle)
}


//delete vehicle

func deleteVehicleByID(c *gin.Context, queries *db.Queries) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid vehicle ID",
        })
        return
    }

    err = queries.DeleteVehicle(c, int32(id))

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to delete vehicle",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Vehicle deleted successfully",
    })
}