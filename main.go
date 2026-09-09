package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Location struct {
	DeviceID   string    `json:"device_id"`
	DeviceName string    `json:"device_name"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	RecordedAt time.Time `json:"recorded_at"`
}

var locations = []Location{}
var pool *pgxpool.Pool
var ctx = context.Background()

func main() {
	r := gin.Default()

	r.GET("/devices", getDevices)
	r.POST("/locations", createLocation)
	r.GET("/locations", getLocations)
	r.GET("/locations/latest", getLatestLocations)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func createLocation(c *gin.Context) {
	var location Location

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	locations = append(locations, location)

	c.JSON(http.StatusOK, location)
}
func getLatestLocations(c *gin.Context) {
	if len(locations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no locations found"})
		return
	}
	c.JSON(http.StatusOK, locations[len(locations)-1])
}
func getLocations(c *gin.Context) {
	if len(locations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no locations found"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func getDevices(c *gin.Context) {
	devices := []string{}
	for i := 0; i != len(locations); i++ {
		devices = append(devices, locations[i].DeviceID)
	}
	c.JSON(http.StatusOK, devices)
}
