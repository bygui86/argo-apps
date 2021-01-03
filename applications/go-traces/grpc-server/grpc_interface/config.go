package grpc_interface

import (
	"github.com/bygui86/go-traces/grpc-server/logging"
	"github.com/bygui86/go-traces/grpc-server/utils"
)

const (
	serverAddressEnvVar = "GRPC_SERVER_ADDRESS"

	serverAddressDefault = "0.0.0.0:50051"
)

func loadConfig() *config {
	logging.Log.Info("Load gRPC configurations")

	return &config{
		address: utils.GetStringEnv(serverAddressEnvVar, serverAddressDefault),
	}
}
