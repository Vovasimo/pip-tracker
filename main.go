package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Location struct {
	DeviceID  string  `json:"device_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var locations = []Location{}

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
