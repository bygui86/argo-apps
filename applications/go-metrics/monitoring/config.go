package monitoring

import (
	"github.com/bygui86/go-metrics/envvars"
	"github.com/bygui86/go-metrics/logging"
)

const (
	// Environment variables -
	monHostEnvVar            = "MONITOR_HOST"
	monPortEnvVar            = "MONITOR_PORT"
	monShutdownTimeoutEnvVar = "MONITOR_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	monHostDefault     = "localhost"
	monPortDefault     = 9090
	monShutdownTimeout = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logging.Log.Debugln("[MONITORING] Setup new Monitoring config...")

	return &Config{
		RestHost:        envvars.GetStringEnv(monHostEnvVar, monHostDefault),
		RestPort:        envvars.GetIntEnv(monPortEnvVar, monPortDefault),
		ShutdownTimeout: envvars.GetIntEnv(monShutdownTimeoutEnvVar, monShutdownTimeout),
	}, nil
}
