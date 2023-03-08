package container

import (
	"math/rand"
	"os"
	"path/filepath"
	"poker/internal/metadata"
	"poker/internal/types"
	"time"
)

// CreateContainer Create rootfs,metadata of new container
func CreateContainer(image, command, name string) (string, error) {
	containerId := generateRandomId(MAX_CONTAINERID)
	imageFilePath := filepath.Join(IMAGE_FOLDER_PATH, image)
	containerPath := filepath.Join(CONTAINER_FOLDER_PATH, containerId)
	rootfsPath := filepath.Join(containerPath, "rootfs")
	MetadataFilePath := filepath.Join(containerPath, "metadata.json")
	ExecFilePath := filepath.Join(containerPath, "exec")

	if name == "" {
		name = generateRandomName()
	}
	if err := checkNameAvailable(name); err != nil {
		return "", err
	}

	// move image to container rootfs
	if err := copyFileOrDir(imageFilePath, rootfsPath); err != nil {
		return "", err
	}

	// move exec to container path
	if err := copyFileOrDir("/var/lib/poker/bin/exec", ExecFilePath); err != nil {
		return "", err
	}

	// write metadata
	if err := metadata.WriteMetadata(MetadataFilePath, makeMetadata(containerId, name, image, command)); err != nil {
		return "", err
	}

	if name != "" {
		nameToContainer[name] = containerId
	}

	return containerId, nil
}

// Generate a new ID, len = n
func generateRandomId(n uint) string {
	b := make([]byte, n)
	length := len(ID_RAND_SOURCE)
	for i := range b {
		b[i] = ID_RAND_SOURCE[rand.Intn(length)]
	}
	return string(b)
}

// Copy File or Directory from src to dst
func copyFileOrDir(src string, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		if err := os.MkdirAll(dst, 0777); err != nil {
			return err
		}
		if list, err := os.ReadDir(src); err == nil {
			for _, item := range list {
				if err = copyFileOrDir(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); err != nil {
					return err
				}
			}
		} else {
			return err
		}
	} else {
		content, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dst, content, 0777); err != nil {
			return err
		}
	}
	return nil
}

// Generate a random name
func generateRandomName() string {
	pre := make([]byte, 8)
	length := len(NAME_RAND_SOURCE)
	for i := range pre {
		pre[i] = NAME_RAND_SOURCE[rand.Intn(length)]
	}
	suf := make([]byte, 5)
	for i := range suf {
		suf[i] = NAME_RAND_SOURCE[rand.Intn(length)]
	}
	return string(pre) + "_" + string(suf)
}

// Make metadata struct
func makeMetadata(id, name, image, command string) *types.ContainerMetadata {
	return &types.ContainerMetadata{
		Id:      id,
		Name:    name,
		Image:   image,
		Created: time.Now(),
		Command: command,
		State:   types.State{Status: "Created"},
	}
}
