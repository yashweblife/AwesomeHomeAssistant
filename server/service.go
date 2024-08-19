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

// it should output the userID or nill
func AuthenticateLogin(Email string, Password string) AuthLoginOutput {
	rows, err := DB.Query("SELECT id, email, password FROM users WHERE email = ? AND password = ?", Email, Password)
	fmt.Println("INSIDE AUTH")
	if err != nil {
		log.Fatal(err)
		return AuthLoginOutput{
			valid:  false,
			userID: "",
		}
	}
	fmt.Println("Passes Query", rows)
	var isEmail string
	var isPassword string
	var userID string

	rows.Scan(&isEmail, &isPassword)
	fmt.Println("Passes Scan")
	if isEmail == Email && isPassword == Password {
		fmt.Println("OUTPUT", isEmail, isPassword)
		return AuthLoginOutput{
			valid:  true,
			userID: userID,
		}
	}
	fmt.Println("OUTPUT", isEmail, isPassword)
	return AuthLoginOutput{
		valid:  false,
		userID: "",
	}
}

func AddUser(Name string, Email string, Password string) string {
	// check if email already exists, get the first instance
	fmt.Println("INSIDE ADD USER")
	rows, err := DB.Exec("SELECT email from users WHERE email = ? LIMIT 1", Email)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	fmt.Println("Passes Query", rows)
	if rows != nil {
		return ""
	}

	Id := uuid.New().String()
	index, err := DB.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", Id, Name, Email, Password)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fmt.Println(index)
	return Id
}

func GetAllUsers() []User {
	var users []User
	fmt.Println("INSIDE GET ALL USERS")
	rows, err := DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Passes Query")
	for rows.Next() {
		fmt.Println("INSIDE NEXT")
		var user User
		err := rows.Scan(&user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
		users = append(users, user)
	}
	fmt.Println("Passes Scan")
	return users
}

func InitUsersDB() {
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
