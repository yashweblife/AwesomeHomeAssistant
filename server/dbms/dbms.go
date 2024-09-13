package dbms

import (
	"database/sql"

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
	// TODO: add a checker for if the user already exists
	// TODO: add a validator for if the user was created
	id := uuid.New().String()
	_, err := DB.Query("INSERT INTO USERS (ID TEXT, EMAIL TEXT, NAME TEXT, PASSWORD, TEXT, DEVICES TEXT) VALUES (?,?,?,?,?)",
		id, email, name, password, "[]")
	if err != nil {
		return "", err
	}
	return id, nil
}

func (d *DBMS) AddDeviceToDB(url, name string) (string, error) {
	// TODO: add checker for if the device already exists
	// TODO: add validator for if the device was created
	id := uuid.New().String()
	_, err := DB.Query("INSERT INTO DEVICES (ID TEXT, URL TEXT, NAME TEXT, COMMANDS TEXT) VALUES (?,?,?,?)", id, url, name, "{}")
	if err != nil {
		return "", err
	}
	return id, nil
}
