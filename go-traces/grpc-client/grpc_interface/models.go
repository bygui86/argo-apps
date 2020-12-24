package grpc_interface

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	config             *config
	ctx                context.Context
	grpcConnection     *grpc.ClientConn
	helloServiceClient HelloServiceClient
	ticker             *time.Ticker
	running            bool
}

type config struct {
	grpcServerAddress     string
	grpcConnectionTimeout time.Duration
	greetingName          string
	greetingInterval      time.Duration
}
