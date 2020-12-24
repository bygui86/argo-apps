package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/kubemq-io/kubemq-go"

	"github.com/bygui86/go-traces/kubemq-consumer/logging"
)

func New(ctx context.Context, consumerName string) (*KubemqConsumer, error) {
	logging.Log.Info("Create new kubemq consumer")

	cfg := loadConfig()

	client, err := kubemq.NewClient(ctx,
		kubemq.WithClientId(consumerName),
		kubemq.WithAddress(cfg.kubemqHost, cfg.kubemqPort),
		// kubemq.WithDefaultChannel(cfg.kubemqChannel),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		return nil, err
	}

	return &KubemqConsumer{
		config:  cfg,
		ctx:     ctx,
		name:    consumerName,
		client:  client,
		running: false,
	}, nil
}

func (c *KubemqConsumer) Start() error {
	logging.Log.Info("Start kubemq consumer")

	if c.client != nil {
		errChan, msgChan, err := c.subscribeToEventStore()
		if err != nil {
			return err
		}

		go c.startConsumer(msgChan, errChan)
		c.running = true
		logging.Log.Info("Kubemq consumer started")
		return nil
	}

	return fmt.Errorf("kubemq consumer start failed: consumer not initialized or already running")
}

func (c *KubemqConsumer) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown kubemq consumer, timeout %d", timeout)

	time.Sleep(time.Duration(timeout) * time.Second)
	c.ctx.Done()
	if c.client != nil {
		err := c.client.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error closing kubemq client: %s", err.Error())
		}
	}
}
