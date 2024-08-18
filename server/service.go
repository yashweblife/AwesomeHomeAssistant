package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *sql.DB

type AuthLoginOutput struct {
	valid  bool
	userID string
}

func AuthenticateLogin(Email string, Password string) AuthLoginOutput {
	rows, err := DB.Query("SELECT id, email, password FROM users WHERE email = ? AND password = ?", Email, Password)
	if err != nil {
		log.Fatal(err)
		return AuthLoginOutput{
			valid:  false,
			userID: "",
		}
	}
	var isEmail string
	var isPassword string
	var userID string

	rows.Scan(&isEmail, &isPassword)
	if isEmail == Email && isPassword == Password {
		return AuthLoginOutput{
			valid:  true,
			userID: userID,
		}
	}
	return AuthLoginOutput{
		valid:  false,
		userID: "",
	}
}

func AddUser(Name string, Email string, Password string) bool {
	// check if email already exists, get the first instance
	rows, err := DB.Exec("SELECT from users WHERE email = ? LIMIT 1", Email)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if rows != nil {
		return false
	}
	
	Id := uuid.New().String()
	index, err := DB.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", Id, Name, Email, Password)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(index)
	return true
}

func GetAllUsers() []User {
	var users []User
	rows, err := DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS users
		 (
			id TEXT, 
			name TEXT, 
			email TEXT, 
			password TEXT, 
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		log.Fatal(err)
	}
}
