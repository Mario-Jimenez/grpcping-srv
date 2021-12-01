package server

import (
	"context"

	"github.com/Mario-Jimenez/grpcping-srv/api/v1/ping"
)

type pingServer struct {
	*ping.UnimplementedPingServer
}

func newPingServer() *pingServer {
	return &pingServer{}
}

func (*pingServer) Ping(context.Context, *ping.PingRequest) (*ping.PingResponse, error) {
	return &ping.PingResponse{}, nil
}
