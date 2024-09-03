package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeviceManager struct {
}

func (dm *DeviceManager) RegisterDevice(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	err := AddDeviceToDB(uuid.New().String(), device.URL, device.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": "Registered"})
}
func (dm *DeviceManager) DeleteDevice(c *gin.Context) {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	err := RemoveDeviceFromDB(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": "Removed"})
}
func (dm *DeviceManager) UpdateDevice(c *gin.Context) {
}
func (dm *DeviceManager) GetDevice(c *gin.Context) {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	var device Device
	err := GetDeviceFromDB(id, &device)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": device})
}
func (dm *DeviceManager) GetDevices(c *gin.Context) {
	var devices []Device
	err := GetAllDevices(&devices)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": devices})
}
func (dm *DeviceManager) SendRequestToDevice(c *gin.Context) {
	type outputData struct {
		Value int `json:"value"`
	}
	res, err := http.Get("http://192.168.0.29:81/value")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
	}
	var data outputData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
	}
	c.JSON(http.StatusOK, data)
}

func (dm *DeviceManager) SendDoesWorkMessage(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Works!"})
}
