package container

import (
	"poker/internal/metadata"
	"poker/pkg/service"
)

func Ps() ([]*service.ContainerInfo, error) {
	metas, err := metadata.ReadAll(CONTAINER_PATH)
	if err != nil {
		return nil, err
	}

	// Convert to service format
	containers := make([]*service.ContainerInfo, len(metas))
	for i, meta := range metas {
		containers[i] = &service.ContainerInfo{
			Id:      meta.Id,
			Name:    meta.Name,
			Image:   meta.Image,
			Created: meta.Created.UnixNano(),
			Command: meta.Command,
			State: &service.State{
				Status: meta.State.Status,
				Pid:    int32(meta.State.Pid),
				Error:  meta.State.Error,
				Start:  meta.State.Start.UnixNano(),
				Finish: meta.State.Finish.UnixNano(),
			},
		}
	}
	return containers, nil
}
