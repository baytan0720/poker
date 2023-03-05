package container

import (
	"poker/internal/metadata"
	"poker/internal/service"
	"syscall"
)

func Stop(containerIds []string) []*service.StartNStopContainerInfo {
	stop := make([]*service.StartNStopContainerInfo, len(containerIds))
	for i, id := range containerIds {
		stop[i] = &service.StartNStopContainerInfo{ContainerId: id}

		// check container id
		containerId, err := findPath(id)
		if err != nil {
			stop[i].Status = 1
			stop[i].Msg = err.Error()
			continue
		}
		containerPath := CONTAINER_FOLDER_PATH + containerId
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
