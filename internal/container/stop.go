package container

import (
	"poker/internal/service"
	"syscall"
)

func Stop(containerIdsOrNames []string) []*service.Answer {
	answers := make([]*service.Answer, len(containerIdsOrNames))
	for i, containerIdOrName := range containerIdsOrNames {
		answers[i] = &service.Answer{ContainerIdOrName: containerIdOrName}
		// check container
		_, _, _, meta, err := checkContainer(containerIdOrName)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		// check container status
		if meta.State.Status == "Exited" {
			continue
		}

		// kill the container
		err = syscall.Kill(meta.State.Pid, syscall.SIGKILL)
		if err != nil {
			answers[i].Status = 1
			answers[i].Msg = err.Error()
			continue
		}

		// need not update metadata.json,
		// when run or start,
		// it will "wait()" cmd exit
	}

	return answers
}
