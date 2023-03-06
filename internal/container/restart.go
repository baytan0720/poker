package container

import "poker/internal/service"

func Restart(containerIds []string) []*service.StartNStopContainerInfo {
	stops := Stop(containerIds)
	starts := Start(containerIds)
	restarts := make([]*service.StartNStopContainerInfo, len(containerIds))
	for i := 0; i < len(containerIds); i++ {
		restarts[i] = &service.StartNStopContainerInfo{ContainerId: containerIds[i]}
		if stops[i].Status != 0 {
			restarts[i].Status = stops[i].Status
			restarts[i].Msg = stops[i].Msg
		}
		if starts[i].Status != 0 {
			restarts[i].Status = starts[i].Status
			restarts[i].Msg = starts[i].Msg
		}
	}
	return restarts
}
