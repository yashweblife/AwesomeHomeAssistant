package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestDBMS(t *testing.T) {
	InitDatabase()
	userID := uuid.New().String()
	t.Run("Add User to DB", func(t *testing.T) {
		err := AddUserToDB(userID, "test", "test@test.test", "0000", nil)
		if err != nil {
			t.Fail()
		}
	})
	t.Run("Get User Info", func(t *testing.T) {
		email := ""
		err := GetUserInfo(userID, &email)
		if err != nil {
			t.Fail()
		}
		t.Log("Email: ", email)
	})
	t.Run("Get All Users", func(t *testing.T) {
		var list []User
		err := GetAllUsers(list)
		if err == false {
			t.Fail()
		}
		t.Log("List: ", list)
	})
	t.Run("Delete User", func(t *testing.T) {
		var didRemoveUser bool
		RemoveUserFromDB(userID, &didRemoveUser)
		if didRemoveUser != true {
			t.Fail()
		}
	})
}
