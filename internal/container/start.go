package container

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"poker/internal/logs"
	"poker/internal/metadata"
	"poker/internal/service"
	"syscall"
	"time"
)

func Start(containerIdsOrNames []string) []*service.Answer {
	answers := make([]*service.Answer, len(containerIdsOrNames))
	for i, containerIdOrName := range containerIdsOrNames {
		answers[i] = &service.Answer{ContainerIdOrName: containerIdOrName}

		_, containerPath, metadataPath, meta, err := checkContainer(containerIdOrName)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		// check container status, if running, ignore it.
		if meta.State.Status == "Running" {
			continue
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

		// bind log to stdout
		logFilePath := filepath.Join(containerPath, "stdout.log")
		f, err := logs.OpenLogs(logFilePath)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}
		cmd.Stdout = f
		cmd.Stderr = f

		// start cmd with non-block
		if err := cmd.Start(); err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		// wait and keep watch on the container state
		go func() {
			// update metadata
			meta.State.Status = "Running"
			meta.State.Start = time.Now()
			meta.State.Pid = cmd.Process.Pid
			meta.State.Error = ""
			if err := metadata.WriteMetadata(metadataPath, meta); err != nil {
				log.Println(err)
			}

			// wait cmd finish
			err := cmd.Wait()
			_ = f.Close()
			meta.State.Finish = time.Now()
			if err != nil {
				meta.State.Error = err.Error()
			}
			meta.State.Status = "Exited"

			if err := metadata.WriteMetadata(metadataPath, meta); err != nil {
				log.Println(err)
			}
		}()
	}

	return answers
}
