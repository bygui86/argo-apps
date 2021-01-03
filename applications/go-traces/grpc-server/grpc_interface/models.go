package grpc_interface

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	config     *config
	listener   net.Listener
	grpcServer *grpc.Server
	running    bool
}

type config struct {
	address string
}
