package rest

import (
	"github.com/bygui86/go-traces/http-client/logging"
	"github.com/bygui86/go-traces/http-client/utils"
)

const (
	restServerHostEnvVar = "REST_SERVER_HOST"
	restServerPortEnvVar = "REST_SERVER_PORT"
	restHostEnvVar       = "REST_HOST"
	restPortEnvVar       = "REST_PORT"

	restServerHostEnvVarDefault = "localhost"
	restServerPortEnvVarDefault = 8080
	restHostEnvVarDefault       = "localhost"
	restPortEnvVarDefault       = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		restServerHost: utils.GetStringEnv(restServerHostEnvVar, restServerHostEnvVarDefault),
		restServerPort: utils.GetIntEnv(restServerPortEnvVar, restServerPortEnvVarDefault),
		restHost:       utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		restPort:       utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
