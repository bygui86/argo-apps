package producer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-traces/kubemq-producer/logging"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (p *KubemqProducer) startProducer() {
	rand.Seed(time.Now().UnixNano())
	p.ticker = time.NewTicker(1 * time.Second)
	counter := 0
	for {
		select {
		case <-p.ticker.C:
			span := opentracing.StartSpan(p.name)

			tags := make(map[string]string, 2)
			tags["sample"] = "sample-value"

			traceErr := opentracing.GlobalTracer().Inject(
				span.Context(),
				opentracing.TextMap,
				opentracing.TextMapCarrier(tags))
			if traceErr != nil {
				logging.SugaredLog.Errorf("Producer failed to inject tracing span: %s", traceErr.Error())
				// continue
			}

			eventStore := p.client.NewEventStore().
				// SetChannel(p.config.kubemqChannel).
				// AddTag("sample", "sample-value").
				SetId(fmt.Sprintf("%s.%d", p.name, counter)).
				SetBody(getMessage(100))
			eventStore.Tags = tags

			// sending stream message
			sendResult, err := eventStore.Send(p.ctx)
			if err != nil {
				logging.SugaredLog.Errorf("Producer failed to send id[%s] msg[%s] tags[%v] to channel %s: %s",
					eventStore.Id, string(eventStore.Body), eventStore.Tags, p.config.kubemqChannel, err.Error())
			}

			// we might have an error in KubeMQ level
			if sendResult.Err != nil {
				logging.SugaredLog.Errorf("Producer didn't send id[%s] msg[%s] tags[%v] to channel %s: %s",
					eventStore.Id, string(eventStore.Body), eventStore.Tags, p.config.kubemqChannel, sendResult.Err.Error())
			} else {
				logging.SugaredLog.Infof("Producer sent successfully id[%s] msg[%s] tags[%v] to channel %s",
					eventStore.Id, string(eventStore.Body), eventStore.Tags, p.config.kubemqChannel)
			}

			counter++

			span.Finish()

		case <-p.ctx.Done():
			return
		}
	}
}

func getMessage(length int64) []byte {
	msg := &Message{
		ID:   uuid.New().String(),
		Data: getRandomString(length),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		logging.SugaredLog.Errorf("JSON marshaling of message %+v failed: %s", msg, err.Error())
	}
	return data
}

func getRandomString(length int64) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
