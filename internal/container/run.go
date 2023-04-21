package container

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"poker/internal/metadata"
	"poker/pkg/pty"
	"strconv"
	"syscall"
	"time"
)

func Run(containerId string) (string, error) {
	containerPath := filepath.Join(CONTAINER_PATH, containerId)
	metadataFilePath := filepath.Join(containerPath, "metadata.json")

	// read metadata
	meta, err := metadata.ReadMetadata(metadataFilePath)
	if err != nil {
		return "", err
	}

	// check container status
	if meta.State.Status == "Running" {
		return "", errors.New("the container is running")
	}

	// isolate namespace
	cmd := exec.Command(filepath.Join(containerPath, "exec"), meta.Command)
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

	// generate a random port for pty
	ttyPort := generateRandomPort()

	l, err := net.Listen("tcp", ":"+ttyPort)
	if err != nil {
		return "", err
	}
	if err := pty.NewTty(l, cmd); err != nil {
		return "", err
	}

	// wait and keep watch on the container state
	go func() {
		// update metadata
		meta.State.Status = "Running"
		meta.State.Start = time.Now()
		meta.State.Pid = cmd.Process.Pid
		if err := metadata.WriteMetadata(metadataFilePath, meta); err != nil {
			log.Println(err)
			err := cmd.Wait()
			if err != nil {
				log.Println(err)
			}
			return
		}

		// wait cmd to exit
		err := cmd.Wait()
		meta.State.Finish = time.Now()
		if err != nil {
			meta.State.Error = err.Error()
		}
		meta.State.Status = "Exited"

		if err := metadata.WriteMetadata(metadataFilePath, meta); err != nil {
			log.Println(err)
		}
	}()

	return ttyPort, nil
}

// generate random port for pty
func generateRandomPort() string {
	return strconv.FormatInt(int64(rand.Intn(10000)+10721), 10)
}
