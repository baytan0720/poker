package engine

import (
	"context"
	"poker/pkg/proto"
)

func (s *Server) Ping(context.Context, *proto.PingReq) (*proto.PongRes, error) {
	return &proto.PongRes{}, nil
}

func (s *Server) Version(context.Context, *proto.VersionReq) (*proto.VersionRes, error) {
	return &proto.VersionRes{}, nil
}

func (s *Server) RunContainer(context.Context, *proto.RunContainerReq) (*proto.RunContainerRes, error) {
	return &proto.RunContainerRes{}, nil
}

func (s *Server) StartContainer(context.Context, *proto.StartContainerReq) (*proto.StartContainerRes, error) {
	return &proto.StartContainerRes{}, nil
}

func (s *Server) StopContainer(context.Context, *proto.StopContainerReq) (*proto.StopContainerRes, error) {
	return &proto.StopContainerRes{}, nil
}

func (s *Server) RestartContainer(context.Context, *proto.RestartContainerReq) (*proto.RestartContainerRes, error) {
	return &proto.RestartContainerRes{}, nil
}

func (s *Server) RemoveContainer(context.Context, *proto.RemoveContainerReq) (*proto.RemoveContainerRes, error) {
	return &proto.RemoveContainerRes{}, nil
}

func (s *Server) ListContainer(context.Context, *proto.ListContainerReq) (*proto.ListContainerRes, error) {
	return &proto.ListContainerRes{}, nil
}

func (s *Server) InspectContainer(context.Context, *proto.InspectContainerReq) (*proto.InspectContainerRes, error) {
	return &proto.InspectContainerRes{}, nil
}

func (s *Server) ExecContainer(ctx context.Context, req *proto.ExecContainerReq) (*proto.ExecContainerRes, error) {
	return &proto.ExecContainerRes{}, nil
}

func (s *Server) LogsContainer(_ context.Context, req *proto.LogsContainerReq) (*proto.LogsContainerRes, error) {
	return &proto.LogsContainerRes{}, nil
}

func (s *Server) RenameContainer(_ context.Context, req *proto.RenameContainerReq) (*proto.RenameContainerRes, error) {
	return &proto.RenameContainerRes{}, nil
}
