package main

import (
	// "encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	InitDatabase()
	auth_route := r.Group("/auth")
	{
		auth_route.POST("login", LoginUser)
		auth_route.POST("register", AddUser)
		auth_route.GET("users", GetUsers)
		auth_route.GET("user", GetUser)
		auth_route.DELETE("user", RemoveUser)
	}
	device_route := r.Group("/device")
	{
		device_route.GET("/", SendDoesWorkMessage)
		device_route.GET("value", SendRequestToDevice)
		device_route.POST("register", RegisterDevice)
		device_route.DELETE("delete", DeleteDevice)
		device_route.GET("devices", GetDevices)
		device_route.GET("device", GetDevice)
	}
	r.Run(PORT)
}
