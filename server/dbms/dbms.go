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
	defer DB.Close()
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT, devices TEXT)")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT, commands TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (d *DBMS) AddUser(name, email, password string) (string, error) {
	userID := uuid.New().String()
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE id = ?", userID).Scan(count)
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", nil
	}
	_, err = DB.Exec("INSERT INTO USERS (id, name, email, password) VALUES (?,?,?,?)", userID, name, email, password)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (d *DBMS) GetUser(userID string) (User, error) {
	var user User
	row := DB.QueryRow("SELECT * FROM USERS WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Devices)
	if err != nil {
		return user, err
	}
	return user, nil
}
