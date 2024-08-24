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
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
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
func AddUserToDB(id, name, email, password string, didCreate *bool) error {
	var count int
	fmt.Println("INSIDE ADD USER TO DB", id, email)
	err := DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE id = ? AND email = ?", id, email).Scan(&count)
	if err != nil {
		return err
	}
	fmt.Println("USER COUNT", count)
	if count > 0 {
		*didCreate = false
		return nil
	}
	_, err = DB.Exec("INSERT INTO USERS (id, name, email, password) VALUES (?,?,?,?)", id, name, email, password)
	if err != nil {
		return err
	}
	var _email string
	err = DB.QueryRow("SELECT email FROM users WHERE id = ?", id).Scan(&_email)
	if err != nil {
		return err
	}
	fmt.Println(_email)
	*didCreate = true
	return nil
}
func GetUserInfo(id string, email *string) {
	err := DB.QueryRow("SELECT email from users WHERE id = ?", id).Scan(&email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(email)
}
func GetAllUsers(list []User) bool {
	var rows *sql.Rows
	var err error

	rows, err = DB.Query("SELECT * from users")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err.Error())
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
