package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"poker/internal/metadata"
	"syscall"
)

// exec command args...
func main() {
	// need command
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	containerPath := os.Args[0][:91]
	metadataPath := filepath.Join(containerPath, "metadata.json")
	command := os.Args[1]
	args := os.Args[2:]
	hostname := "container"
	rootfsPath := filepath.Join(containerPath, "rootfs")

	// read metadata
	if meta, err := metadata.ReadMetadata(metadataPath); err == nil {
		hostname = meta.Name
	}

	// init
	if err := syscall.Sethostname([]byte(hostname)); err != nil {
		panic(err)
	}
	if err := syscall.Chroot(rootfsPath); err != nil {
		panic(err)
	}
	if err := syscall.Chdir("/"); err != nil {
		panic(err)
	}
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		panic(err)
	}
	path, err := exec.LookPath(command)
	if err == nil {
		command = path
	}
	if err := syscall.Exec(command, args, os.Environ()); err != nil {
		panic(err)
	}
}
