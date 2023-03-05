package container

import (
	"errors"
	"path/filepath"
)

const (
	ID_RAND_SOURCE        = "abcdefghijklmnopqrstuvwxyz0123456789"
	NAME_RAND_SOURCE      = "abcdefghijklmnopqrstuvwxyz"
	MAX_CONTAINERID       = 64
	IMAGE_FOLDER_PATH     = "/var/lib/poker/images/"
	CONTAINER_FOLDER_PATH = "/var/lib/poker/containers/"
)

// find container id
func find(containerId string) (string, error) {
	if len(containerId) == MAX_CONTAINERID {
		return containerId, nil
	}
	pattern := CONTAINER_FOLDER_PATH + containerId + "*"
	matchs, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}
	if len(matchs) < 1 {
		return "", errors.New("the container does not exist")
	}
	if len(matchs) > 1 {
		return "", errors.New("the matching container num > 1")
	}
	return matchs[0], nil
}
