package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Mario-Jimenez/grpcping-srv/internal/server"
	"google.golang.org/grpc"
)

type serverRunner interface {
	Run(chan<- error) error
	Shutdown() error
}

type Service struct {
	server serverRunner
}

func New() (*Service, error) {
	s := &Service{}
	err := s.setupServer()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) setupServer() error {
	var opts []grpc.ServerOption
	s.server = server.NewGRPCServer(":53001", opts...)

	return nil
}

func (s *Service) Run() {
	var exit chan error
	err := s.server.Run(exit)
	if err != nil {
		// TODO: log error
		fmt.Println(err)
		s.shutdown()
		return
	}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-exit:
		// TODO: log error
		fmt.Println(err)
		s.shutdown()
		os.Exit(1)
	case <-osSignal:
		s.shutdown()
	}
}

func (s *Service) shutdown() {
	err := s.server.Shutdown()
	if err != nil {
		// TODO: log error
		fmt.Println(err)
	}
}
