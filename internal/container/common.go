package container

import (
	"errors"
	"fmt"
	"path/filepath"
	"poker/internal/metadata"
)

const (
	ID_RAND_SOURCE        = "abcdefghijklmnopqrstuvwxyz0123456789"
	NAME_RAND_SOURCE      = "abcdefghijklmnopqrstuvwxyz"
	MAX_CONTAINERID       = 64
	IMAGE_FOLDER_PATH     = "/var/lib/poker/images/"
	CONTAINER_FOLDER_PATH = "/var/lib/poker/containers/"
)

var nameToContainer = map[string]string{}

func init() {
	metas, err := metadata.ReadAll(CONTAINER_FOLDER_PATH)
	if err != nil {
		panic(err)
	}
	for _, meta := range metas {
		nameToContainer[meta.Name] = meta.Id
	}
}

// find container path
func findPath(containerId string) (string, error) {
	if len(containerId) == MAX_CONTAINERID {
		return CONTAINER_FOLDER_PATH + containerId, nil
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

func checkNameAvailable(name string) error {
	if name == "" {
		return nil
	}
	if len(name) > 16 {
		return errors.New("name is too long")
	}
	if tmp, ok := nameToContainer[name]; ok {
		return errors.New(fmt.Sprintf("The container name \"%s\" is already in use by container %s", name, tmp))
	}
	return nil
}

func checkName(containerIdOrName string) string {
	// if length = MAX_CONTAINERID, it must be container id
	if len(containerIdOrName) == MAX_CONTAINERID {
		return containerIdOrName
	}

	// if find container id by name, return container id
	if id, ok := nameToContainer[containerIdOrName]; ok {
		return id
	}
	return containerIdOrName
}
