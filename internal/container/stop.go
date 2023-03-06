package container

import (
	"poker/internal/metadata"
	"poker/internal/service"
	"syscall"
)

func Stop(containerIds []string) []*service.StartNStopContainerInfo {
	stop := make([]*service.StartNStopContainerInfo, len(containerIds))
	for i, containerId := range containerIds {
		stop[i] = &service.StartNStopContainerInfo{ContainerId: containerId}

		// check container id
		containerPath, err := findPath(containerId)
		if err != nil {
			stop[i].Status = 1
			stop[i].Msg = err.Error()
			continue
		}
		metadataFilePath := containerPath + "/metadata.json"

		// read metadata
		meta, err := metadata.ReadMetadata(metadataFilePath)
		if err != nil {
			stop[i].Status = 1
			stop[i].Msg = err.Error()
			continue
		}

		// check container status
		if meta.State.Status == "Exited" {
			continue
		}

		// kill the container
		err = syscall.Kill(meta.State.Pid, syscall.SIGKILL)
		if err != nil {
			stop[i].Status = 1
			stop[i].Msg = err.Error()
			continue
		}
	}

	return stop
}
