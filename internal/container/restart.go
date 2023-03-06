package container

import "poker/internal/service"

func Restart(containerIdsOrNames []string) []*service.StartNStopContainerInfo {
	stops := Stop(containerIdsOrNames)
	starts := Start(containerIdsOrNames)
	restarts := make([]*service.StartNStopContainerInfo, len(containerIdsOrNames))
	for i := 0; i < len(containerIdsOrNames); i++ {
		restarts[i] = &service.StartNStopContainerInfo{ContainerIdOrName: containerIdsOrNames[i]}
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
