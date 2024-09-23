package dbms

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DBMS struct {
	DB *sql.DB
}

func (d *DBMS) Init() error {
	var err error
	d.DB, err = sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		return err
	}
	_, err = d.DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT, devices TEXT)")
	if err != nil {
		return err
	}
	_, err = d.DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT, commands TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (d *DBMS) AddUser(name, email, password string) (string, error) {
	userID := uuid.New().String()
	var count int
	err := d.DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE email = ?", email).Scan(&count)
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("User already exists")
	}
	_, err = d.DB.Exec("INSERT INTO USERS (id, name, email, password, devices) VALUES (?,?,?,?,?)", userID, name, email, password, "{}")
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (d *DBMS) GetUser(userID string) (User, error) {
	var user User
	row := d.DB.QueryRow("SELECT * FROM USERS WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Devices)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d *DBMS) RemoveUser(userID string) error {
	_, err := d.DB.Exec("DELETE FROM USERS WHERE id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
func (d *DBMS) EditUser() {}

func (d *DBMS) AddDevice(url, name string) (string, error) {
	deviceID := uuid.New().String()
	_, err := d.DB.Exec("INSERT INTO DEVICES (id, url, name) VALUES (?,?,?)", deviceID, url, name)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (d *DBMS) GetDevice(id string) (Device, error) {}

func (d *DBMS) RemoveDevice(id string) error {}
