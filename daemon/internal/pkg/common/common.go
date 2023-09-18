package common

import (
	"math/rand"
	"net"
	"os"
	"path/filepath"
)

const (
	ContainerIDLength       = 64
	containerNameLength     = 10
	containerIDRandSource   = "abcdefghijklmnopqrstuvwxyz0123456789"
	containerNameRandSource = "abcdefghijklmnopqrstuvwxyz"
)

// GenerateRandomId Generate a new ID
func GenerateRandomId() string {
	b := make([]byte, ContainerIDLength)
	length := len(containerIDRandSource)
	for i := range b {
		b[i] = containerIDRandSource[rand.Intn(length)]
	}
	return string(b)
}

// GenerateRandomName Generate a new name
func GenerateRandomName() string {
	b := make([]byte, containerNameLength)
	length := len(containerNameRandSource)
	for i := range b {
		if i == containerNameLength/2 {
			b[i] = '_'
			continue
		}
		b[i] = containerNameRandSource[rand.Intn(length)]
	}
	return string(b)
}

// CopyFileOrDir Copy File or Directory from src to dst
func CopyFileOrDir(src string, dst string) error {
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
				if err = CopyFileOrDir(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); err != nil {
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

func ListenOnUnix(socketPath string) (net.Listener, error) {
	_ = os.Remove(socketPath)

	laddr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		return nil, err
	}

	l, err := net.ListenUnix("unix", laddr)
	if err != nil {
		return nil, err
	}

	return l, nil
}
