package dbms

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DBMS struct {
}

var DB *sql.DB

func (d *DBMS) Init() error {
	DB, err := sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT)")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (d *DBMS) AddUserToDB(email, password, name string) (string, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE EMAIL= ? ", email).Scan(&count)
	if count != 0 {
		return "", errors.New("Email Already exists")
	}

	// TODO: add a validator for if the user was created

	id := uuid.New().String()
	_, err = DB.Query("INSERT INTO USERS (ID TEXT, EMAIL TEXT, NAME TEXT, PASSWORD, TEXT, DEVICES TEXT) VALUES (?,?,?,?,?)",
		id, email, name, password, "[]")
	if err != nil {
		return "", err
	}
	type User struct {
		name     string
		id       string
		email    string
		password string
		devices  string
	}
	var user User
	err = DB.QueryRow("SELECT * FROM USERS WHERE ID = ?", id).Scan(&user.id, &user.name, &user.email, &user.password, &user.devices)
	if err != nil {
		return "", err
	}
	if user.email != email {
		return "", errors.New("Failed to create user")
	}
	return id, nil
}

func (d *DBMS) AddDeviceToDB(user_id, url, name string) (string, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM DEVICES WHERE URL = ?", url).Scan(&count)
	if err != nil {
		return "", err
	}
	if count != 0 {
		return "", errors.New("Device already exists")
	}
	// TODO: add validator for if the device was created
	id := uuid.New().String()
	_, err = DB.Exec("INSERT INTO DEVICES (ID TEXT, URL TEXT, NAME TEXT, COMMANDS TEXT) VALUES (?,?,?,?)", id, url, name, "{}")
	if err != nil {
		return "", err
	}
	type Device struct {
		id       string
		url      string
		name     string
		commands string
	}
	var device Device
	err = DB.QueryRow("SELECT * FROM DEVICES WHERE ID = ?", id).Scan(&device.id, &device.url, &device.name, &device.commands)
	if err != nil {
		return "", err
	}
	if device.url != url {
		return "", errors.New("Failed to create device")
	}
	_, err = DB.Query("UPDATE USERS SET DEVICES = JSON_ARRAY_APPEND(DEVICES, '$', ?) WHERE ID = ?", id, user_id)
	return id, nil
}

func (d *DBMS) ValidateUser(email, password string) (string, error) {

	return "", nil
}
