package iot

type IOT struct {
}

func (iot *IOT) Init() error {
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
