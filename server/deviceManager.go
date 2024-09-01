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

func (dm *DeviceManager) RegisterDevice(c *gin.Context) error {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	err := AddDeviceToDB(uuid.New().String(), device.URL, device.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": "Registered"})
	return nil
}
func (dm *DeviceManager) DeleteDevice(c *gin.Context) error {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	err := RemoveDeviceFromDB(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": "Removed"})
	return nil
}
func (dm *DeviceManager) UpdateDevice(c *gin.Context) error {
	return nil
}
func (dm *DeviceManager) GetDevice(c *gin.Context) error {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	var device Device
	err := GetDeviceFromDB(id, &device)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": device})
	return nil
}
func (dm *DeviceManager) GetDevices(c *gin.Context) error {
	var devices []Device
	err := GetAllDevices(&devices)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": devices})
	return nil
}
func (dm *DeviceManager) SendRequestToDevice(c *gin.Context) error {
	type outputData struct {
		Value int `json:"value"`
	}
	res, err := http.Get("http://192.168.0.29:81/value")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
		return err
	}
	var data outputData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNoContent, gin.H{"data": "No Host"})
		return err
	}
	c.JSON(http.StatusOK, data)
	return nil
}

func (dm *DeviceManager) SendDoesWorkMessage(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Works!"})
}
