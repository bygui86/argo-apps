package grpc_interface

import (
	"time"

	"github.com/bygui86/go-traces/grpc-client/logging"
	"github.com/bygui86/go-traces/grpc-client/utils"
)

const (
	serverAddressEnvVar     = "GRPC_SERVER_ADDRESS"
	connectionTimeoutEnvVar = "GRPC_CONNECTION_TIMEOUT" // in seconds
	greetingNameEnvVar      = "GRPC_GREETING_NAME"
	greetingIntervalEnvVar  = "GRPC_GREETING_INTERVAL" // in seconds

	serverAddressDefault     = "0.0.0.0:50051"
	connectionTimeoutDefault = 2
	greetingNameDefault      = "ANONYMOUS"
	greetingIntervalDefault  = 1
)

func loadConfig() *config {
	logging.Log.Info("Load gRPC configurations")

	return &config{
		grpcServerAddress:     utils.GetStringEnv(serverAddressEnvVar, serverAddressDefault),
		grpcConnectionTimeout: time.Duration(utils.GetIntEnv(connectionTimeoutEnvVar, connectionTimeoutDefault)) * time.Second,
		greetingName:          utils.GetStringEnv(greetingNameEnvVar, greetingNameDefault),
		greetingInterval:      time.Duration(utils.GetIntEnv(greetingIntervalEnvVar, greetingIntervalDefault)) * time.Second,
	}
}
