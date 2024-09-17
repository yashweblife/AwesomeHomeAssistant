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

}
