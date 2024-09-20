package dbms

import (
	"testing"
)

func TestDBMS(t *testing.T) {
	dbms := DBMS{}
	var user User
	t.Run("Init", func(t *testing.T) {
		err := dbms.Init()
		if err != nil {
			t.Fail()
		}
	})
	t.Run("Add User", func(t *testing.T) {
		userID, err := dbms.AddUser("test", "test", "test")
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		user.ID = userID
		user.Name = "test"
		user.Email = "test"
		user.Password = "test"
		user.Devices = "{}"
		t.Log("Added User", user)
	})
	t.Run("Get User", func(t *testing.T) {
		user, err := dbms.GetUser(user.ID)
		if err != nil {
			t.Fail()
		}
		t.Log("Got User", user)
	})
}
