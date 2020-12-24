package consumer

import (
	"context"

	"github.com/kubemq-io/kubemq-go"
)

type KubemqConsumer struct {
	config  *Config
	ctx     context.Context
	name    string
	client  *kubemq.Client
	running bool
}

type Config struct {
	kubemqHost    string
	kubemqPort    int
	kubemqChannel string
	kubemqGroup   string
}

type Message struct {
	ID   string
	Data string
}
