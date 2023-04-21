package metadata

import "time"

type Container struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Image   string    `json:"image"`
	Created time.Time `json:"created"`
	Command string    `json:"command"`
	State   State     `json:"state"`
}

type State struct {
	Status string    `json:"status"` // Running/Exited
	Pid    int       `json:"pid"`
	Error  string    `json:"error"`
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
}
