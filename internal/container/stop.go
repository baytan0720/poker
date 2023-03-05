package container

import (
	"poker/internal/metadata"
	"poker/internal/service"
	"syscall"
	"time"
)

func Stop(containerIds []string) []*service.StartNStopContainerInfo {
	stop := make([]*service.StartNStopContainerInfo, len(containerIds))
	for i, containerId := range containerIds {
		stop[i] = &service.StartNStopContainerInfo{ContainerId: containerId}
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

		// update metadata
		meta.State.Finish = time.Now()
		meta.State.Status = "Exited"
		_ = metadata.WriteMetadata(metadataFilePath, meta)
	}

	return stop
}
