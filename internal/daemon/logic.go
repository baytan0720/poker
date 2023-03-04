package daemon

import (
	"context"
	"poker/internal/container"
	"poker/internal/service"
)

func (d *Daemon) Ping(context.Context, *service.PingReq) (*service.PingRes, error) {
	return &service.PingRes{}, nil
}

func (d *Daemon) CreateContainer(_ context.Context, req *service.CreateContainerReq) (*service.CreateContainerRes, error) {
	res := &service.CreateContainerRes{}

	id, err := container.CreateContainer(req.Image, req.Command, req.Name)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	res.ContainerId = id

	return res, nil
}

func (d *Daemon) StartContainer(ctx context.Context, req *service.StartContainerReq) (*service.StartContainerRes, error) {
	res := &service.StartContainerRes{}
	err := container.Run(req.ContainerId)
	if err != nil {
		res.Status = 1
		res.Msg = err.Error()
	}
	return res, nil
}
