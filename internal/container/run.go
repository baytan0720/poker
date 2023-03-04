package container

import (
	"errors"
	"os"
	"os/exec"
	"poker/internal/metadata"
	"poker/internal/types"
	"syscall"
	"time"
)

func Run(containerId string) error {
	containerPath := CONTAINER_FOLDER_PATH + containerId
	MetadataFilePath := containerPath + "/metadata.json"

	// read metadata
	meta, err := metadata.ReadMetadata(MetadataFilePath)
	if err != nil {
		return err
	}
	if meta.State.Status == "running" {
		return errors.New("container is already running")
	}

	// isolate namespace
	cmd := exec.Command(containerPath+"/exec", meta.Command)
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

	// bind log to stdout
	logFilePath := CONTAINER_FOLDER_PATH + containerId + "/stdout.log"
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	cmd.Stdout = f

	// start cmd with non-block
	if err := cmd.Start(); err != nil {
		return err
	}

	// wait and keep watch on the container state
	go func() {
		err := cmd.Wait()
		if err != nil {
			meta.State = types.State{
				Status: "Exited",
				Pid:    "",
				Error:  err.Error(),
				Start:  time.Time{},
				Finish: time.Now(),
			}
		}
		_ = f.Close()
	}()

	return nil
}
