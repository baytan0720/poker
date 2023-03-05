package main

import (
	"os"
	"os/exec"
	"poker/internal/metadata"
	"syscall"
)

// exec command args...
func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	containerPath := os.Args[0][:91]
	metadataPath := containerPath + "metadata.json"

	// update metadata
	meta, err := metadata.ReadMetadata(metadataPath)
	if err != nil {
		panic(err)
	}

	// init
	command := os.Args[1]
	args := os.Args[2:]
	rootfsPath := containerPath + "rootfs"
	if err := syscall.Sethostname([]byte(meta.Name)); err != nil {
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
	if err != nil {
		panic(err)
	}
	if err := syscall.Exec(path, args, os.Environ()); err != nil {
		panic(err)
	}
}
