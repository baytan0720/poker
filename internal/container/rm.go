package container

import (
	"os"
	"poker/pkg/service"
)

func Remove(containerIdsOrNames []string) []*service.Answer {
	answers := make([]*service.Answer, len(containerIdsOrNames))
	for i, containerIdOrName := range containerIdsOrNames {
		answers[i] = &service.Answer{ContainerIdOrName: containerIdOrName}
		_, containerPath, _, meta, err := checkContainer(containerIdOrName)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		if meta.State.Status == "running" {
			answers[i].Status = 1
			answers[i].Msg = "stop the container before remove"
			continue
		}

		err = os.RemoveAll(containerPath)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		delete(nameToContainer, meta.Name)
	}
	return answers
}
