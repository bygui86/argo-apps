package consumer

import (
	"encoding/json"

	"github.com/kubemq-io/kubemq-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/bygui86/go-traces/kubemq-consumer/logging"
)

func (c *KubemqConsumer) subscribeToEventStore() (chan error, <-chan *kubemq.EventStoreReceive, error) {
	// subscribe to event-store channel
	errChan := make(chan error, 1)
	msgChan, err := c.client.SubscribeToEventsStore(
		c.ctx,
		c.config.kubemqChannel,
		c.config.kubemqGroup,
		errChan,
		kubemq.StartFromNewEvents())
	if err != nil {
		return nil, nil, err
	}

	return errChan, msgChan, nil
}

func (c *KubemqConsumer) startConsumer(msgChan <-chan *kubemq.EventStoreReceive, errChan chan error) {
	for {
		select {
		case streamMsg := <-msgChan:
			spanContext, spanErr := opentracing.GlobalTracer().Extract(
				opentracing.TextMap,
				opentracing.TextMapCarrier(streamMsg.Tags))
			if spanErr != nil {
				logging.SugaredLog.Errorf("Error extracting span from tags: %s", spanErr.Error())
			}
			span := opentracing.StartSpan(c.name, ext.RPCServerOption(spanContext))

			msg := &Message{}
			jsonErr := json.Unmarshal(streamMsg.Body, msg)
			if jsonErr == nil {
				logging.SugaredLog.Infof("Message received from channel %s:\n    timestamp[%v], clientId[%s], id[%s], sequence[%d], \n    metadata[%s], tags[%+v], message[%+v]",
					streamMsg.Channel, streamMsg.Timestamp, streamMsg.ClientId, streamMsg.Id, streamMsg.Sequence,
					streamMsg.Metadata, streamMsg.Tags, msg)
			} else {
				logging.SugaredLog.Errorf("JSON unmarshal of body failed: %s", jsonErr.Error())
				logging.SugaredLog.Infof("Message received from channel %s:\n    timestamp[%v], clientId[%s], id[%s], sequence[%d], \n    metadata[%s], tags[%+v], body[%s]",
					streamMsg.Channel, streamMsg.Timestamp, streamMsg.ClientId, streamMsg.Id, streamMsg.Sequence,
					streamMsg.Metadata, streamMsg.Tags, string(streamMsg.Body))
			}

			span.Finish()

		case err := <-errChan:
			logging.SugaredLog.Errorf("Error received from channel %s: %s", c.config.kubemqChannel, err.Error())

		case <-c.ctx.Done():
			return
		}
	}
}
