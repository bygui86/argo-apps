package producer

import (
	"context"
	"time"

	"github.com/kubemq-io/kubemq-go"
)

type KubemqProducer struct {
	config  *Config
	ctx     context.Context
	name    string
	client  *kubemq.Client
	ticker  *time.Ticker
	running bool
}

type Config struct {
	kubemqHost    string
	kubemqPort    int
	kubemqChannel string
}

type Message struct {
	ID   string
	Data string
}
