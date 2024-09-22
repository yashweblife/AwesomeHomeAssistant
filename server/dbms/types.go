package dbms

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Devices  string
}

type Device struct {
	ID       string
	URL      string
	Name     string
	Commands string
}
