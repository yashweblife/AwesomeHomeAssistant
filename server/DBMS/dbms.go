package dbms

import (
	"database/sql"
)

type DBMS struct {
	DB *sql.DB
}

func (dbms *DBMS) Initialize() error {

	var err error
	dbms.DB, err = sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		return err
	}
	_, err = dbms.DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT)")
	if err != nil {
		return err
	}
	_, err = dbms.DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT, owner TEXT)")
	if err != nil {
		return err
	}
	_, err = dbms.DB.Exec("CREATE TABLE IF NOT EXISTS GROUPS (id TEXT, name TEXT, owner TEXT, members TABLE)")
	if err != nil {
		return err
	}
	return nil
}
