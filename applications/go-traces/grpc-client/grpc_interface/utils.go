package grpc_interface

import (
	"context"
	"time"

	"github.com/bygui86/go-traces/grpc-client/logging"
)

func (c *Client) startGreeting() {
	logging.SugaredLog.Infof("Start greeting %s every %s",
		c.config.greetingName, c.config.greetingInterval.String())
	c.ticker = time.NewTicker(c.config.greetingInterval)

	for {
		select {
		case <-c.ticker.C:
			// WARN: the connection context is one-shot, it must be refreshed before every request
			connectionCtx, cancel := context.WithTimeout(context.Background(), c.config.grpcConnectionTimeout)
			defer cancel()
			response, err := c.helloServiceClient.SayHello(connectionCtx, &HelloRequest{Name: c.config.greetingName})
			if err != nil {
				logging.SugaredLog.Errorf("Greeting %s failed: %v", c.config.greetingName, err.Error())
				continue
			}
			logging.SugaredLog.Infof(response.Greeting)

		case <-c.ctx.Done():
			return
		}
	}
}
