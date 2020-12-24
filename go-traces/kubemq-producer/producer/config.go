package producer

import (
	"github.com/bygui86/go-traces/kubemq-producer/logging"
	"github.com/bygui86/go-traces/kubemq-producer/utils"
)

const (
	kubemqHostEnvVar    = "KUBEMQ_HOST"
	kubemqPortEnvVar    = "KUBEMQ_PORT"
	kubemqChannelEnvVar = "KUBEMQ_CHANNEL"

	kubemqHostDefault    = "localhost"
	kubemqPortDefault    = 50000
	kubemqChannelDefault = "stream.channel"
)

func loadConfig() *Config {
	logging.Log.Debug("Load kubemq producer configurations")
	return &Config{
		kubemqHost:    utils.GetStringEnv(kubemqHostEnvVar, kubemqHostDefault),
		kubemqPort:    utils.GetIntEnv(kubemqPortEnvVar, kubemqPortDefault),
		kubemqChannel: utils.GetStringEnv(kubemqChannelEnvVar, kubemqChannelDefault),
	}
}
