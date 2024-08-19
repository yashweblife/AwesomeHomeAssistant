package main

import (
	"log"
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
		log.Println("Error binding request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	if auth.Email == "" || auth.Password == "" {
		log.Println("Error: Email and password are required")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Email and password are required"})
		return
	}
	log.Println("Received request for login: ", auth)
	test := AuthenticateLogin(auth.Email, auth.Password)
	if test == nil {
		log.Println("Error: Authentication returned null")
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Authentication returned null"})
		return
	}
	log.Println("Response from authentication: ", test)
	c.JSON(http.StatusOK, gin.H{"data": test})
}
func AddUserToDB(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		log.Println("Error binding request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	log.Println("Received request for adding user: ", user)

	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Println("Error: Name, email, and password are required")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Name, email, and password are required"})
		return
	}

	Id := AddUser(user.Name, user.Email, user.Password)

	log.Println("Response from adding user: ", Id)
	c.JSON(http.StatusOK, gin.H{"data": &Id})
}
func GetUsers(c *gin.Context) {
	users := GetAllUsers()
	c.JSON(200, gin.H{"data": users})
}
func main() {
	r := gin.Default()
	InitUsersDB()
	auth_route := r.Group("/auth")
	{
		auth_route.POST("login", LoginUser)
		auth_route.POST("register", AddUserToDB)
		auth_route.GET("users", GetUsers)
	}
	r.Run(PORT)
}
