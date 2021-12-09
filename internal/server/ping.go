package server

import (
	"context"
	"fmt"

	"github.com/Mario-Jimenez/grpcping-srv/api/v1/ping"
)

type pingServer struct {
	*ping.UnimplementedPingServer
}

func newPingServer() *pingServer {
	return &pingServer{}
}

func (*pingServer) Ping(context.Context, *ping.PingRequest) (*ping.PingResponse, error) {
	fmt.Println("Ping received")
	return &ping.PingResponse{}, nil
}
