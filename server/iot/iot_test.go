package iot

import (
	"testing"
)

func TestIOT(t *testing.T) {
	iot := IOT{}
	t.Run("Init", func(t *testing.T) {
		err := iot.Init()
		if err != nil {
			t.Fail()
		}
	})

	t.Run("Check If Online", func(t *testing.T) {
		online, error := iot.CheckIfOnline()
		if error != nil {
			t.Fail()
		}
		if online {
			t.Log("Working")
		}
	})
	t.Run("Get Commands", func(t *testing.T) {
		commands, error := iot.GetCommands()
		if error != nil {
			t.Fail()
		}
		if commands == "{}" {
			t.Log("Working")
		}
	})
	t.Run("Call Command", func(t *testing.T) {
		command, error := iot.CallCommand("test")
		if error != nil {
			t.Fail()
		}
		if command == "{}" {
			t.Log("Working")
		}
	})

}
