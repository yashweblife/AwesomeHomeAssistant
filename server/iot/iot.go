package iot

import (
	"errors"
	"net/http"
)

type IOT struct {
	url string
}

func (iot *IOT) Init() error {
	if iot.url == "" {
		iot.url = "http://192.168.0.20:81/"
	}
	client := http.DefaultClient
	req, err := http.NewRequest("GET", iot.url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("Failed to connect to IOT")
	}
	return nil
}
func (iot *IOT) CheckIfOnline() (bool, error) {
	return true, nil
}
func (iot *IOT) GetCommands() (string, error) {
	return "{}", nil
}
func (iot *IOT) CallCommand(name string) (string, error) {
	return "{}", nil
}
