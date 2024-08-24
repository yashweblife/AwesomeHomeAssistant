package main

type Command struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

type Device struct {
	ID       string    `json:"id"`
	URL      string    `json:"url"`
	Type     bool      `json:"type"`
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}
type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Devices  []Device `json:"devices"`
}
