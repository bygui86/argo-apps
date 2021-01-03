package producer

import (
	"context"
	"time"

	"github.com/kubemq-io/kubemq-go"

	"github.com/bygui86/go-traces/kubemq-producer/logging"
)

func New(ctx context.Context, producerName string) (*KubemqProducer, error) {
	logging.Log.Info("Create new kubemq producer")

	cfg := loadConfig()

	client, err := kubemq.NewClient(ctx,
		kubemq.WithClientId(producerName),
		kubemq.WithAddress(cfg.kubemqHost, cfg.kubemqPort),
		kubemq.WithDefaultChannel(cfg.kubemqChannel),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		return nil, err
	}

	return &KubemqProducer{
		config:  cfg,
		ctx:     ctx,
		name:    producerName,
		client:  client,
		running: false,
	}, nil
}

func (p *KubemqProducer) Start() {
	logging.Log.Info("Start kubemq producer")

	if p.client != nil {
		go p.startProducer()
		p.running = true
		logging.Log.Info("Kubemq producer started")
		return
	}

	logging.Log.Error("Kubemq producer start failed: producer not initialized or already running")
}

func (p *KubemqProducer) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown kubemq producer, timeout %d", timeout)

	p.ticker.Stop()
	p.ctx.Done()
	time.Sleep(time.Duration(timeout) * time.Second)
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error closing kubemq client: %s", err.Error())
		}
	}
}
