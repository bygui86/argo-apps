package producer

import (
	"github.com/bygui86/go-traces/kafka-producer/logging"
	"github.com/bygui86/go-traces/kafka-producer/utils"
)

const (
	kafkaBootstrapServersEnvVar = "KAFKA_BOOTSTRAP_SERVERS" // comma-separated list host or host:port
	kafkaPortEnvVar             = "KAFKA_PORT"
	kafkaTopicEnvVar            = "KAFKA_TOPIC"

	kafkaBootstrapServersDefault = "localhost"
	kafkaPortDefault             = 9092
	kafkaTopicDefault            = "my-topic"
)

func loadConfig() *Config {
	logging.Log.Debug("Load kafka producer configurations")
	return &Config{
		kafkaBootstrapServers: utils.GetStringEnv(kafkaBootstrapServersEnvVar, kafkaBootstrapServersDefault),
		kafkaPort:             utils.GetIntEnv(kafkaPortEnvVar, kafkaPortDefault),
		kafkaTopic:            utils.GetStringEnv(kafkaTopicEnvVar, kafkaTopicDefault),
	}
}
