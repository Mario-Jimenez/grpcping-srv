package server

import (
	"errors"
	"net"

	"github.com/Mario-Jimenez/grpcping-srv/api/v1/ping"
	"google.golang.org/grpc"
)

type grpcServer struct {
	server  *grpc.Server
	address string
	port    string
}

func NewGRPCServer(port string, grpcOpts ...grpc.ServerOption) *grpcServer {
	gsrv := grpc.NewServer(grpcOpts...)

	pingServer := newPingServer()
	ping.RegisterPingServer(gsrv, pingServer)

	return &grpcServer{server: gsrv, port: port}
}

func (s *grpcServer) Run(exit chan<- error) error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return errors.New(err.Error())
	}

	s.address = listener.Addr().String()

	go func() {
		if err := s.server.Serve(listener); err != nil {
			exit <- err
		}
	}()

	return nil
}

func (s *grpcServer) Shutdown() error {
	s.server.GracefulStop()

	return nil
}
