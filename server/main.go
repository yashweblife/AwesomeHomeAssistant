package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	dm := DeviceManager{}
	um := UserManager{}
	InitDatabase()
	auth_route := r.Group("/auth")
	{
		auth_route.POST("login", um.LoginUser)
		auth_route.POST("register", um.AddUser)
		auth_route.GET("users", um.GetUsers)
		auth_route.GET("user", um.GetUser)
		auth_route.DELETE("user", um.RemoveUser)
	}
	device_route := r.Group("/device")
	{
		device_route.GET("/", dm.SendDoesWorkMessage)
		device_route.GET("value", dm.SendRequestToDevice)
		device_route.POST("register", dm.RegisterDevice)
		device_route.DELETE("delete", dm.DeleteDevice)
		device_route.GET("devices", dm.GetDevices)
		device_route.GET("device", dm.GetDevice)
	}
	r.Run(PORT)
}
