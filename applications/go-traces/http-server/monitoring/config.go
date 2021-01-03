package monitoring

import (
	"github.com/bygui86/go-traces/http-server/logging"
	"github.com/bygui86/go-traces/http-server/utils"
)

const (
	monitorHostEnvVar = "MONITOR_HOST"
	monitorPortEnvVar = "MONITOR_PORT"

	monitorHostDefault = "localhost"
	monitorPortDefault = 9090
)

func loadConfig() *Config {
	logging.Log.Debug("Load monitoring configurations")
	return &Config{
		restHost: utils.GetStringEnv(monitorHostEnvVar, monitorHostDefault),
		restPort: utils.GetIntEnv(monitorPortEnvVar, monitorPortDefault),
	}
}
