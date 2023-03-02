package container

import (
	"os"
	"os/exec"
	"poker/alert"
	"syscall"
)

func CreateRuncontainer(containerID string, args []string) {
	rootPath := "/var/lib/poker/containers/" + containerID + "/rootfs"
	if len(args) < 1 {
		os.Exit(0)
	}
	if err := syscall.Sethostname([]byte(containerID[:10])); err != nil {
		alert.Fatal(err.Error())
	}
	if err := syscall.Chroot(rootPath); err != nil {
		alert.Fatal(err.Error())
	}
	if err := syscall.Chdir("/"); err != nil {
		alert.Fatal(err.Error())
	}
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		alert.Fatal(err.Error())
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		alert.Fatal(err.Error())
	}
	syscall.Exec(path, args, os.Environ())
}
