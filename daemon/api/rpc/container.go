package rpc

import (
	"encoding/json"
	gin "github.com/baytan0720/gin-grpc"
	"poker/daemon/api/rpc_util"
	"poker/daemon/internal/container"
	"poker/daemon/internal/manager"
	"poker/daemon/middleware"
	"poker/pkg/errno"
	"poker/pkg/proto"
)

type ContainerRpc struct{}

func (r *ContainerRpc) Functions() rpc_util.Functions {
	return rpc_util.Functions{
		{
			Name:    "RunContainer",
			Handler: r.RunContainer,
			Hooks:   gin.HandlersChain{middleware.ValidName},
		},
		{
			Name:    "StartContainer",
			Handler: r.StartContainer,
		},
		{
			Name:    "StopContainer",
			Handler: r.StopContainer,
		},
		{
			Name:    "RestartContainer",
			Handler: r.RestartContainer,
		},
		{
			Name:    "RemoveContainer",
			Handler: r.RemoveContainer,
		},
		{
			Name:    "ListContainer",
			Handler: r.ListContainer,
		},
		{
			Name:    "InspectContainer",
			Handler: r.InspectContainer,
		},
		{
			Name:    "ExecContainer",
			Handler: r.ExecContainer,
		},
		{
			Name:    "LogsContainer",
			Handler: r.LogsContainer,
		},
		{
			Name:    "RenameContainer",
			Handler: r.RenameContainer,
		},
	}
}

func (r *ContainerRpc) RunContainer(c *gin.Context) {
	req := &proto.RunContainerReq{}
	c.BindRequest(req)

	image, err := manager.GetImage(req.Image)
	if err != nil {
		c.Response(&proto.RunContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	ctr := container.NewContainer(image, genCreateOptions(req)...)

	if err := ctr.Init(); err != nil {
		c.Response(&proto.RunContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	ctr.Start()

	manager.AddContainer(ctr)

	c.Response(&proto.RunContainerRes{
		ErrorCode:  errno.OK,
		Id:         ctr.Id,
		SocketPath: ctr.Config.TtySocket,
	})
}

func (r *ContainerRpc) StartContainer(c *gin.Context) {
	req := &proto.StartContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.StartContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	ctr.Start()
}

func (r *ContainerRpc) StopContainer(c *gin.Context) {
	req := &proto.StopContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.StopContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	ctr.Stop()
}

func (r *ContainerRpc) RestartContainer(c *gin.Context) {
	req := &proto.RestartContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.RestartContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	ctr.Restart()
}

func (r *ContainerRpc) RemoveContainer(c *gin.Context) {
	req := &proto.RemoveContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.RemoveContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	err = ctr.Remove(req.Force)
	if err != nil {
		c.Response(&proto.RemoveContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	manager.RemoveContainer(ctr)
}

func (r *ContainerRpc) ListContainer(c *gin.Context) {
	containers := manager.ListContainers()
	b, err := json.Marshal(containers)
	if err != nil {
		c.Response(&proto.ListContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}
	c.Response(&proto.ListContainerRes{
		Containers: b,
	})
}

func (r *ContainerRpc) InspectContainer(c *gin.Context) {
	req := &proto.InspectContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.InspectContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}
	b, err := json.Marshal(ctr)
	if err != nil {
		c.Response(&proto.InspectContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}
	c.Response(&proto.InspectContainerRes{
		Container: b,
	})
}

func (r *ContainerRpc) ExecContainer(c *gin.Context) {

}

func (r *ContainerRpc) LogsContainer(c *gin.Context) {
	req := &proto.LogsContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.LogsContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}
	logs := ctr.Logs()
	c.Response(&proto.LogsContainerRes{
		Logs: logs,
	})
}

func (r *ContainerRpc) RenameContainer(c *gin.Context) {
	req := &proto.RenameContainerReq{}
	c.BindRequest(req)
	ctr, err := manager.GetContainer(req.IdOrName)
	if err != nil {
		c.Response(&proto.RenameContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	if err := manager.ValidateName(req.NewName); err != nil {
		c.Response(&proto.RenameContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}

	manager.RemoveContainer(ctr)
	defer manager.AddContainer(ctr)

	if err := ctr.Rename(req.NewName); err != nil {
		c.Response(&proto.RenameContainerRes{
			ErrorCode: errno.Errno(err),
		})
		return
	}
}

func genCreateOptions(req *proto.RunContainerReq) []container.CreateOption {
	var options []container.CreateOption

	if req.Name != "" {
		options = append(options, container.WithName(req.Name))
	}
	if req.Tty {
		options = append(options, container.WithTty())
	}
	if req.Command != "" {
		options = append(options, container.WithCommandArgs(req.Command, req.Args))
	}
	if len(req.Env) > 0 {
		options = append(options, container.WithEnv(req.Env))
	}
	if len(req.ExposedPorts) > 0 {
		options = append(options, container.WithExposedPorts(req.ExposedPorts))
	}
	if len(req.Volumes) > 0 {
		options = append(options, container.WithVolumes(req.Volumes))
	}
	if req.Hostname != "" {
		options = append(options, container.WithHostname(req.Hostname))
	}
	if req.User != "" {
		options = append(options, container.WithUser(req.User))
	}
	if req.WorkingDir != "" {
		options = append(options, container.WithWorkingDir(req.WorkingDir))
	}
	if req.Restart {
		options = append(options, container.WithAutoRestart())
	}
	if req.Rm {
		options = append(options, container.WithAutoRemove())
	}

	return options
}
