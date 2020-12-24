package consumer

import (
	"github.com/bygui86/go-traces/kubemq-consumer/logging"
	"github.com/bygui86/go-traces/kubemq-consumer/utils"
)

const (
	kubemqHostEnvVar    = "KUBEMQ_HOST"
	kubemqPortEnvVar    = "KUBEMQ_PORT"
	kubemqChannelEnvVar = "KUBEMQ_CHANNEL"
	kubemqGroupEnvVar   = "KUBEMQ_CONSUMER_GROUP"

	kubemqHostDefault    = "localhost"
	kubemqPortDefault    = 50000
	kubemqChannelDefault = "stream.channel"
	kubemqGroupDefault   = "my-group"
)

func loadConfig() *Config {
	logging.Log.Debug("Load kubemq consumer configurations")
	return &Config{
		kubemqHost:    utils.GetStringEnv(kubemqHostEnvVar, kubemqHostDefault),
		kubemqPort:    utils.GetIntEnv(kubemqPortEnvVar, kubemqPortDefault),
		kubemqChannel: utils.GetStringEnv(kubemqChannelEnvVar, kubemqChannelDefault),
		kubemqGroup:   utils.GetStringEnv(kubemqGroupEnvVar, kubemqGroupDefault),
	}
}
