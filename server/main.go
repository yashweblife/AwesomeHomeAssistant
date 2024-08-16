package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

type AuthLoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var auth AuthLoginType
	if err := c.Bind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	var test = AuthenticateLogin(auth.Email, auth.Password)
	c.JSON(200, gin.H{"data": test})
}
func AddUserToDB(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	fmt.Println(user)
	AddUser(user.Name, user.Email, user.Password)
	c.JSON(200, gin.H{"data": "1"})
}
func GetUsers(c *gin.Context) {
	users := GetAllUsers()
	c.JSON(200, gin.H{"data": users})
}
func main() {
	r := gin.Default()
	InitDB()
	auth_route := r.Group("/auth")
	{
		auth_route.POST("login", LoginUser)
		auth_route.POST("register", AddUserToDB)
		auth_route.GET("users", GetUsers)
	}
	r.Run(PORT)
}
