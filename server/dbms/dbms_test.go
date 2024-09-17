package dbms

import (
	"testing"
)

func TestDBMS(t *testing.T){
	dbms := DBMS{}
	t.Run("Init", func (t *testing.T) {
		err := dbms.Init()
		if err != nil {
			t.Errorf("Init() error = %v", err)
		}

		t.Log("Init() ok")
	})
}