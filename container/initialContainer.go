package container

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"poker/alert"
	"strings"
	"syscall"
	"time"
)

// CreateInitialcontainer for isolation the original system
func CreateInitialcontainer(image string, isInteractive, isTty, isDetach bool, args ...string) {
	// Make container rootfs
	containerID := generatecontainerId(MAX_CONTAINERID)
	imageFilePath := IMAGE_FOLDER_PATH + image
	containerPath := CONTAINER_FOLDER_PATH + containerID + "/rootfs"
	if err := copyFileOrDirectory(imageFilePath, containerPath); err != nil {
		alert.Fatal(err.Error())
	}
	if isDetach {
		isTty = false
		fmt.Println(containerID)
	}
	var cmd *exec.Cmd
	if len(args) == 0 {
		cmd = exec.Command("/proc/self/exe", "init", containerID)
	} else {
		cmd = exec.Command("/proc/self/exe", "init", containerID, strings.Join(args, " "))
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	if isTty {
		if isInteractive {
			cmd.Stdin = os.Stdin
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		logFilePath := CONTAINER_FOLDER_PATH + containerID + "/stdout.log"
		f, err := os.Create(logFilePath)
		if err != nil {
			alert.Fatal(err.Error())
		}
		cmd.Stdout = f
	}
	if err := cmd.Start(); err != nil {
		alert.Fatal(err.Error())
	}
	if !isDetach {
		cmd.Wait()
	}
}

// copy image to container
func copyFileOrDirectory(src string, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		if err := os.MkdirAll(dst, 0777); err != nil {
			return err
		}
		if list, err := ioutil.ReadDir(src); err == nil {
			for _, item := range list {
				if err = copyFileOrDirectory(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); err != nil {
					return err
				}
			}
		} else {
			return err
		}
	} else {
		content, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(dst, content, 0777); err != nil {
			return err
		}
	}
	return nil
}

func generatecontainerId(n uint) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}
	return string(b)
}
