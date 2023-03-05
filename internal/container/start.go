package container

import (
	"log"
	"os"
	"os/exec"
	"poker/internal/metadata"
	"poker/internal/service"
	"syscall"
	"time"
)

func Start(containerIds []string) []*service.StartNStopContainerInfo {
	start := make([]*service.StartNStopContainerInfo, len(containerIds))
	for i, containerId := range containerIds {
		start[i] = &service.StartNStopContainerInfo{ContainerId: containerId}
		containerPath := CONTAINER_FOLDER_PATH + containerId
		metadataFilePath := containerPath + "/metadata.json"

		// read metadata
		meta, err := metadata.ReadMetadata(metadataFilePath)
		if err != nil {
			start[i].Status = 1
			start[i].Msg = err.Error()
			continue
		}

		// check container status
		if meta.State.Status == "Running" {
			continue
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
			start[i].Status = 1
			start[i].Msg = err.Error()
			continue
		}
		cmd.Stdout = f
		cmd.Stderr = f

		// start cmd with non-block
		if err := cmd.Start(); err != nil {
			start[i].Status = 1
			start[i].Msg = err.Error()
			continue
		}

		// wait and keep watch on the container state
		go func() {
			// update metadata
			meta.State.Status = "Running"
			meta.State.Start = time.Now()
			meta.State.Pid = cmd.Process.Pid
			if err := metadata.WriteMetadata(metadataFilePath, meta); err != nil {
				log.Println(err)
			}

			// wait cmd exit
			err := cmd.Wait()
			_ = f.Close()

			meta.State.Finish = time.Now()

			if err != nil {
				meta.State.Error = err.Error()
			}
			meta.State.Status = "Exited"

			if err := metadata.WriteMetadata(metadataFilePath, meta); err != nil {
				log.Println(err)
				return
			}
		}()
	}

	return start
}
