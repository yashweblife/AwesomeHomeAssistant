package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoginUser(c *gin.Context) error {
	type AuthLoginType struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var auth AuthLoginType
	if err := c.Bind(&auth); err != nil {
		log.Println("Error binding request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	if auth.Email == "" || auth.Password == "" {
		log.Println("Error: Email and password are required")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Email and password are required"})
		return errors.New("Email and password are required")
	}
	log.Println("Received request for login: ", auth)
	err := AuthenticateLoginAttempt(auth.Email, auth.Password)
	if err != nil {
		log.Println("Error: Authentication returned null")
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Authentication returned null"})
		return errors.New("Authentication returned null")
	}
	log.Println("Response from authentication: ", err)
	c.JSON(http.StatusOK, gin.H{"data": "Success"})
	return nil
}
func AddUser(c *gin.Context) error {
	var user User
	userID := uuid.New().String()
	fmt.Println("User ID: ", userID)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	var didCreateUser bool
	err := AddUserToDB(userID, user.Name, user.Email, user.Password, &didCreateUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	fmt.Println("User created: ", didCreateUser)
	if !didCreateUser {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return errors.New("Bad Request")
	}
	c.JSON(http.StatusOK, gin.H{"data": userID})
	return nil
}
func GetUsers(c *gin.Context) error {
	var users []User
	err := GetAllUsers(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": users})
	return nil
}
func GetUser(c *gin.Context) error {
	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	var userEmail string
	err := GetUserInfo(id.Id, &userEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": userEmail})
	return nil
}
func EditUserInfo(c *gin.Context) error {

	type input struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": "Edited"})
	return nil
}
func RemoveUser(c *gin.Context) error {

	type input struct {
		Id string `json:"id"`
	}
	var id input
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	var didRemoveUser bool
	err := RemoveUserFromDB(id.Id, &didRemoveUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return err
	}
	c.JSON(200, gin.H{"data": didRemoveUser})
	return nil
}
