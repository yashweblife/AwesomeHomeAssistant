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
	Valid  bool
	UserID string
}

func AuthenticateLogin(email string, password string) AuthLoginOutput {
	var userID string
	var dbEmail string
	var dbPassword string

	err := DB.QueryRow("SELECT id, email, password FROM users WHERE email = ? AND password = ?", email, password).Scan(&userID, &dbEmail, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return AuthLoginOutput{
				Valid:  false,
				UserID: "",
			}
		}
		log.Println("Error:", err)
		return AuthLoginOutput{
			Valid:  false,
			UserID: "",
		}
	}

	if email == dbEmail && password == dbPassword {
		return AuthLoginOutput{
			Valid:  true,
			UserID: userID,
		}
	}

	return AuthLoginOutput{
		Valid:  false,
		UserID: "",
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
