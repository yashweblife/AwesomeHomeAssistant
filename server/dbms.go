package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT, devices JSONB) ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT, type TEXT, commands JSONB) ")
	if err != nil {
		log.Fatal(err)
	}
}

// Authentication
func AddUserToDB(id, name, email, password string) bool {
	//check if user exists by email
	var err error
	var result *sql.Rows
	fmt.Println(id, name, email, password)
	result, err = DB.Query("SELECT email from users WHERE email = ? LIMIT 1", email)
	for result.Next() {
		var e string
		result.Scan(&e)
		fmt.Println(e)
	}

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println("Inserting user: ", id)
	_, err = DB.Query("INSERT INTO users (id, name, email, password, devices) VALUES (?, ?, ?, ?, ?)", id, name, email, password, []Device{})
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	fmt.Println("Inserted user: ", id)
	return true

}
func GetUserInfo(id string, email *string) {
	var rows *sql.Rows
	var err error

	rows, err = DB.Query("SELECT * from users WHERE id = ? LIMIT 1", id)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Println(rows)
}
func GetAllUsers(list []User) bool {
	var rows *sql.Rows
	var err error

	rows, err = DB.Query("SELECT * from users")
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Errorf(err.Error())
			return false
		}
		list = append(list, user)
	}
	fmt.Println(list)
	return (true)
}
func RemoveUser(id string) {}
func RemoveAllUsers()      {}
func AuthenticateLoginAttempt(email, password string) bool {
	return true
}

// Devices
func AddDeviceToDB(url, name, id, TYPE string, commands []Command) {}
func AddDevcieToUserDeviceList(id string)                          {}
func GetDeviceInfo(id string)                                      {}
func GetAllDevices()                                               {}
func RemoveDevice(id string)                                       {}
func RemoveAllDevices()                                            {}
