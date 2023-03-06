package daemon

import (
	"context"
	"log"
	"poker/internal/container"
	"poker/internal/service"
	"strings"
)

func (d *Daemon) Ping(context.Context, *service.PingReq) (*service.PingRes, error) {
	log.Println("ping")
	return &service.PingRes{}, nil
}

func (d *Daemon) CreateContainer(_ context.Context, req *service.CreateContainerReq) (*service.CreateContainerRes, error) {
	log.Println("create container from", req.Image)
	res := &service.CreateContainerRes{}

	id, err := container.CreateContainer(req.Image, req.Command, req.Name)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
		log.Println(err)
	}
	res.ContainerId = id

	return res, nil
}

func (d *Daemon) RunContainer(_ context.Context, req *service.RunContainerReq) (*service.RunContainerRes, error) {
	log.Println("run container", req.ContainerId)
	res := &service.RunContainerRes{}
	ptyPort, err := container.Run(req.ContainerId)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	res.PtyPort = ptyPort
	return res, nil
}

func (d *Daemon) StartContainer(_ context.Context, req *service.StartContainersReq) (*service.StartContainersRes, error) {
	log.Println("start containers", strings.Join(req.ContainerIdsOrNames, " "))
	return &service.StartContainersRes{
		StartNStopContainerInfo: container.Start(req.ContainerIdsOrNames),
	}, nil
}

func (d *Daemon) StopContainer(_ context.Context, req *service.StopContainersReq) (*service.StopContainersRes, error) {
	log.Println("start containers", strings.Join(req.ContainerIdsOrNames, " "))
	return &service.StopContainersRes{
		StartNStopContainerInfo: container.Stop(req.ContainerIdsOrNames),
	}, nil
}

func (d *Daemon) RestartContainer(_ context.Context, req *service.RestartContainersReq) (*service.RestartContainersRes, error) {
	log.Println("restart containers", req.ContainerIdsOrNames)
	return &service.RestartContainersRes{
		StartNStopContainerInfo: container.Restart(req.ContainerIdsOrNames),
	}, nil
}

func (d *Daemon) PsContainer(context.Context, *service.PsContainersReq) (*service.PsContainersRes, error) {
	res := &service.PsContainersRes{}
	log.Println("ps containers")
	containers, err := container.Ps()
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	res.Containers = containers
	return res, nil
}

func (d *Daemon) LogsContainer(_ context.Context, req *service.LogsContainerReq) (*service.LogsContainerRes, error) {
	res := &service.LogsContainerRes{}
	log.Println("logs containers", req.ContainerIdOrName)
	logs, err := container.Logs(req.ContainerIdOrName)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	res.Logs = logs
	return res, nil
}

func (d *Daemon) RenameContainer(_ context.Context, req *service.RenameContainerReq) (*service.RenameContainerRes, error) {
	res := &service.RenameContainerRes{}
	log.Println("rename", req.NewName, "container", req.ContainerIdOrName)
	err := container.Rename(req.ContainerIdOrName, req.NewName)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	return res, nil
}
