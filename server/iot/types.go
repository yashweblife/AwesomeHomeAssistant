package iot

type DeviceCommand struct {
	Name string `json:"name"`
	Info string `json:"info"`
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
