package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	// "encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	test := AuthenticateLoginAttempt(auth.Email, auth.Password)
	if !test {
		log.Println("Error: Authentication returned null")
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Authentication returned null"})
		return
	}
	log.Println("Response from authentication: ", test)
	c.JSON(http.StatusOK, gin.H{"data": test})
}
func AddUser(c *gin.Context) {
	var user User
	userID := uuid.New().String()
	fmt.Println("User ID: ", userID)
	if err := c.Bind(&user); err != nil {
		fmt.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	var didCreateUser bool
	err := AddUserToDB(userID, user.Name, user.Email, user.Password, &didCreateUser)
	if err != nil {
		fmt.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	fmt.Println("User created: ", didCreateUser)
	if !didCreateUser {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userID})
}
func GetUsers(c *gin.Context) {
	var users []User
	valid := GetAllUsers(users)
	fmt.Println(valid)
	c.JSON(200, gin.H{"data": users})
}
func GetUser(c *gin.Context) {
	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	var userEmail string
	GetUserInfo(id.Id, &userEmail)
	fmt.Println("userEmail: ", userEmail)
	c.JSON(200, gin.H{"data": userEmail})
}
func EditUserInfo(c *gin.Context) {

	type input struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	c.JSON(200, gin.H{"data": "Edited"})
}
func RemoveUser(c *gin.Context) {

	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	var didRemoveUser bool
	RemoveUserFromDB(id.Id, &didRemoveUser)
	c.JSON(200, gin.H{"data": didRemoveUser})
}

func SendRequestToDevice(c *gin.Context) {

	var url string
	if err := c.Bind(&url); err != nil {
		log.Println("Error binding request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}
	fmt.Println("Received request for device: ", url)
	res, err := http.Get(url + "/value")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": string(body)})
}

func SendDoesWorkMessage(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Works!"})
}

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
		device_route.POST("/", SendDoesWorkMessage)
		device_route.GET("value", SendDoesWorkMessage)
	}
	r.Run(PORT)
}
