package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserManager struct {
}

func (um *UserManager) LoginUser(c *gin.Context) {
	type AuthLoginType struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var auth AuthLoginType
	if err := c.Bind(&auth); err != nil {
		log.Println("Error binding request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	if auth.Email == "" || auth.Password == "" {
		log.Println("Error: Email and password are required")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Email and password are required"})
	}
	log.Println("Received request for login: ", auth)
	err := AuthenticateLoginAttempt(auth.Email, auth.Password)
	if err != nil {
		log.Println("Error: Authentication returned null")
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Authentication returned null"})
	}
	log.Println("Response from authentication: ", err)
	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}
func (um *UserManager) AddUser(c *gin.Context) {
	var user User
	userID := uuid.New().String()
	fmt.Println("User ID: ", userID)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	var didCreateUser bool
	err := AddUserToDB(userID, user.Name, user.Email, user.Password, &didCreateUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	fmt.Println("User created: ", didCreateUser)
	if !didCreateUser {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(http.StatusOK, gin.H{"data": userID})
}
func (um *UserManager) GetUsers(c *gin.Context) {
	var users []User
	err := GetAllUsers(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": users})
}
func (um *UserManager) GetUser(c *gin.Context) {
	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	var userEmail string
	err := GetUserInfo(id.Id, &userEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": userEmail})
}
func (um *UserManager) EditUserInfo(c *gin.Context) {

	type input struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": "Edited"})
}
func (um *UserManager) RemoveUser(c *gin.Context) {

	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	var didRemoveUser bool
	err := RemoveUserFromDB(id.Id, &didRemoveUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
	}
	c.JSON(200, gin.H{"data": didRemoveUser})
}
