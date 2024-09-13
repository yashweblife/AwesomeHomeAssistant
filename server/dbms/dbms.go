package dbms

import (
	"database/sql"

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
