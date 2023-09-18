package container

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	now := time.Now()
	containers := ContainerSlice{
		&Container{CreatedAt: now.Add(-20)},
		&Container{CreatedAt: now.Add(-40)},
		&Container{CreatedAt: now.Add(-10)},
		&Container{CreatedAt: now.Add(-30)},
	}

	sort.Sort(containers)
	assert.Equal(t, containers[0].CreatedAt, now.Add(-10))
	assert.Equal(t, containers[1].CreatedAt, now.Add(-20))
	assert.Equal(t, containers[2].CreatedAt, now.Add(-30))
	assert.Equal(t, containers[3].CreatedAt, now.Add(-40))
}
