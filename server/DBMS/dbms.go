package dbms

import (
	"database/sql"
)

var userScheme = `
CREATE TABLE IF NOT EXIST USERS
(
id TEXT,
name TEXT,
email TEXT,
password TEXT
)
`

var deviceScheme = `
CREATE TABLE IF NOT EXIST DEVICES
(
ID TEXT
NAME TEXT
URL TEXT
OWNER TEXT
)
`

func Initialize() error {

	DB, err := sql.Open("sqlite3", "./AwesomeHA.db")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS USERS (id TEXT, name TEXT, email TEXT, password TEXT)")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS DEVICES (id TEXT, url TEXT, name TEXT, owner TEXT)")
	if err != nil {
		return err
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS GROUPS (id TEXT, name TEXT, owner TEXT, members TABLE)")
	if err != nil {
		return err
	}
	return nil
}

func AddUserToDB() {}

func RemoveUserFromDB() {}

func CheckIfUserIsValid() {}

func GetUserFromDB() {}

func GetAllUsers() {}

func AddDeviceToDB() {}

func RemoveDeviceFromDB() {}
