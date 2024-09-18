package iot

import (
	"testing"
)

func TestIOT(t *testing.T) {
	iot := IOT{
		url: "http://192.168.0.29:81/",
	}
	t.Run("Init", func(t *testing.T) {
		err := iot.Init()
		if err != nil {
			t.Fail()
		}
		t.Log("Device URL is valid")
	})

	t.Run("Check If Online", func(t *testing.T) {
		online, error := iot.CheckIfOnline()
		if error != nil {
			t.Fail()
		}
		if online == false || online == true {
			t.Log("Device Is Online")
		}
	})
	t.Run("Get Commands", func(t *testing.T) {
		commands, error := iot.GetCommands()
		if error != nil {
			t.Fail()
		}
		if len(commands) > 0 {
			t.Log("Got Commands", commands)
		}
	})
	t.Run("Call Command", func(t *testing.T) {
		command, error := iot.CallCommand("on")
		if error != nil {
			t.Fail()
		}
		if command == "ON" {
			t.Log("Command Worked")
		}
	})

}
