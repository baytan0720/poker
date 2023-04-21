package main

import (
	"poker/pkg/service"
	"testing"
)

func TestPs(t *testing.T) {
	printTitle()
	info := &service.ContainerInfo{
		Id:      "asgdasgnkasdnwfaaf",
		Name:    "mysql",
		Image:   "base",
		Created: 1678017868982543000,
		Command: "bash",
		State: &service.State{
			Status: "Exited",
			Pid:    3667,
			Start:  1678017868982543000,
			Finish: 1678019890721453000,
			Error:  "kill -9",
		},
	}
	printPs(info)
	printDetailTitle()
	printPsDetail(info)
}
