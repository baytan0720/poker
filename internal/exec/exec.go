package main

import (
	"os"
	"os/exec"
	"poker/internal/metadata"
	"strconv"
	"syscall"
	"time"
)

// exec command args...
func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	// update metadata
	meta, err := metadata.ReadMetadata("./metadata.json")
	if err != nil {
		panic(err)
	}
	meta.State.Status = "Running"
	meta.State.Pid = strconv.FormatInt(int64(os.Getpid()), 10)
	meta.State.Start = time.Now()
	if err := metadata.WriteMetadata("./metadata.json", meta); err != nil {
		panic(err)
	}

	// init
	command := os.Args[1]
	args := os.Args[2:]
	rootPath := "./rootfs"
	if err := syscall.Sethostname([]byte("container")); err != nil {
		panic(err)
	}
	if err := syscall.Chroot(rootPath); err != nil {
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
