package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterDevice(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	AddDeviceToDB(uuid.New().String(), device.URL, device.Name)
	fmt.Println(device)
	c.JSON(200, gin.H{"data": "Registered"})
}
func DeleteDevice(c *gin.Context) {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	RemoveDeviceFromDB(id)
	c.JSON(200, gin.H{"data": "Removed"})
}
func UpdateDevice(c *gin.Context) {
}
func GetDevice(c *gin.Context) {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	var device Device
	GetDeviceFromDB(id, &device)
	c.JSON(200, gin.H{"data": device})
}
func GetDevices(c *gin.Context) {
	var devices []Device
	GetAllDevices(&devices)
	c.JSON(200, gin.H{"data": devices})
}
func SendRequestToDevice(c *gin.Context) {
	type outputData struct {
		Value int `json:"value"`
	}
	res, err := http.Get("http://192.168.0.29:81/value")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data outputData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, data)
}

func SendDoesWorkMessage(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Works!"})
}
