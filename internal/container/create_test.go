package container

import "testing"

func TestCreateContainer(t *testing.T) {
	id, err := CreateContainer("base", "/bin/bash", "test")
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}
