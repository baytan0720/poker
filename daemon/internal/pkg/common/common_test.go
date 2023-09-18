package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomId(t *testing.T) {
	id := GenerateRandomId()
	assert.Len(t, id, ContainerIDLength)
	t.Log(id)
}

func TestGenerateRandomName(t *testing.T) {
	name := GenerateRandomName()
	assert.Len(t, name, containerNameLength)
	t.Log(name)
}
