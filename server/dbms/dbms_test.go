package dbms

import (
	"testing"
)

func TestDBMS(t *testing.T) {
	dbms := DBMS{}
	t.Run("Init", func(t *testing.T) {
		err := dbms.Init()
		if err != nil {
			t.Fail()
		}
	})
	t.Run("AddUserToDB", func(t *testing.T) {
		id, err := dbms.AddUserToDB("test", "test", "test")
		if err != nil {
			t.Fail()
		}
		if id == "" {
			t.Fail()
		}
		t.Log(id)
	})
}
