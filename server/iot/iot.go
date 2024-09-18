package iot

import (
	"encoding/json"
	"errors"
	"io"
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
	if iot.url == "" {
		iot.url = "http://192.168.0.20:81/"
	}
	client := http.DefaultClient
	req, err := http.NewRequest("GET", iot.url, nil)
	if err != nil {
		return false, err
	}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

type DeviceCommand struct {
	Name string `json:"name"`
	info string `json:"info"`
}
type DeviceInfo struct {
	Type     string          `json:"type"`
	Info     string          `json:"info"`
	IP       string          `json:"ip"`
	Commands []DeviceCommand `json:"commands"`
}
type DeviceResponse struct {
	Data string `json:"data"`
}

func (iot *IOT) GetCommands() ([]DeviceCommand, error) {
	if iot.url == "" {
		iot.url = "http://192.168.0.20:81/"
	}
	client := http.DefaultClient
	req, err := http.NewRequest("GET", iot.url, nil)
	if err != nil {
		return []DeviceCommand{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []DeviceCommand{}, err
	}
	if res.StatusCode != 200 {
		return []DeviceCommand{}, errors.New("Failed to connect to IOT")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []DeviceCommand{}, err
	}
	var data DeviceInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []DeviceCommand{}, err
	}
	return data.Commands, nil
}
func (iot *IOT) CallCommand(name string) (string, error) {
	if iot.url == "" {
		iot.url = "http://192.168.0.20:81/"
	}
	client := http.DefaultClient
	req, err := http.NewRequest("POST", iot.url+name, nil)
	if err != nil {
		return "{}", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "{}", err
	}
	if res.StatusCode != 200 {
		return "{}", errors.New("Failed to connect to IOT")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "{}", err
	}
	var data DeviceResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "{}", err
	}
	return data.Data, nil
}
