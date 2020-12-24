package rest

import (
	"github.com/bygui86/go-metrics/envvars"
	"github.com/bygui86/go-metrics/logging"
)

const (
	// Environment variables -
	restHostEnvVar            = "REST_HOST"
	restPortEnvVar            = "REST_PORT"
	restShutdownTimeoutEnvVar = "REST_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault            = "localhost"
	restPortDefault            = 8080
	restShutdownTimeoutDefault = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logging.Log.Debugln("[REST] Setup new REST server config...")

	return &Config{
		RestHost:        envvars.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:        envvars.GetIntEnv(restPortEnvVar, restPortDefault),
		ShutdownTimeout: envvars.GetIntEnv(restShutdownTimeoutEnvVar, restShutdownTimeoutDefault),
	}, nil
}
