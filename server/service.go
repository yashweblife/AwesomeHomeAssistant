package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *sql.DB

func AuthenticateLogin(Email string, Password string) bool {
	rows, err := DB.Query("SELECT email, password FROM users WHERE email = ? AND password = ?", Email, Password)
	if err != nil {
		log.Fatal(err)
	}

	var isEmail string
	var isPassword string

	rows.Scan(&isEmail, &isPassword)
	if isEmail == Email && isPassword == Password {
		return true
	}
	return false
}

func AddUser(Name string, Email string, Password string) bool {

	index, err := DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", Name, Email, Password)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(index)
	return true
}

func GetAllUsers() []User {
	var users []User
	rows, err := DB.Query("SELECT name, email, password FROM users")
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
