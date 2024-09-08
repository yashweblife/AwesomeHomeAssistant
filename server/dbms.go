package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() error {
	fmt.Println("Connecting to database...")
	DB, err := sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	fmt.Print("Connected to database\nCreating tables:\tUSERS\tDEVICES\n")
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("USERS table created")
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("DEVICES table created")
	return nil
}

// Authentication
func AddUserToDB(id, name, email, password string, didCreate *bool) error {
	var count int
	fmt.Println("INSIDE ADD USER TO DB", id, email)
	err := DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE email = ?", email).Scan(&count)
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
func GetUserInfo(id string, email *string) error {
	var output string
	err := DB.QueryRow("SELECT email from users WHERE id = ?", id).Scan(&output)
	if err != nil {
		return err
	}
	*email = output
	return nil
}
func GetAllUsers(list []User) error {
	rows, err := DB.Query("SELECT * from users")
	if err != nil {
		return err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return err
		}
		list = append(list, user)
	}
	return nil
}
func EditUser(id, name, email, password string) error {

	return nil
}
func RemoveUserFromDB(id string, didRemoveUser *bool) error {
	stmt, err := DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		*didRemoveUser = false
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		*didRemoveUser = false
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		*didRemoveUser = false
		return err
	}

	if affected == 0 {
		*didRemoveUser = false
	}
	*didRemoveUser = true
	return nil
}
func RemoveAllUsers() error {
	stmt, err := DB.Prepare("DELETE * FROM users")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func AuthenticateLoginAttempt(email, password string) error {
	return nil
}

// Devices
func AddDeviceToDB(id, url, name string) error {
	_, err := DB.Exec("INSERT INTO DEVICES (id, url, name) VALUES (?,?,?)", id, url, name)
	if err != nil {
		return err
	}

	return nil
}
func GetDeviceFromDB(id string, device *Device) error {

	return nil
}
func GetAllDevices(list *[]Device) error {

	*list = []Device{}

	rows, err := DB.Query("SELECT * from devices")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var device Device
		err := rows.Scan(&device.ID, &device.URL, &device.Name)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		*list = append(*list, device)
	}
	fmt.Println(*list)

	return nil
}
func RemoveDeviceFromDB(id string) error {

	return nil
}
func RemoveAllDevices() error {

	return nil
}
