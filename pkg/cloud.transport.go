package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jfyne/live"
	"gocloud.dev/pubsub"
)

type Transport struct {
	topic *pubsub.Topic
	debug bool
}

func newTransport(ctx context.Context, debug bool) (*Transport, error) {
	topic, err := pubsub.OpenTopic(ctx, "mem://broadcast")
	if err != nil {
		return nil, err
	}

	return &Transport{topic: topic, debug: debug}, nil
}

func (c *Transport) Publish(ctx context.Context, topic string, msg live.Event) error {
	if msg.Data == nil {
		if raw, err := json.Marshal(msg.SelfData); nil != err {
			return err
		} else {
			msg.Data = raw
		}
	}

	data, err := json.Marshal(live.TransportMessage{Topic: topic, Msg: msg})

	if c.debug {
		fmt.Println(">> ", string(data))
	}

	if err != nil {
		return fmt.Errorf("could not publish event: %w", err)
	}

	return c.topic.Send(ctx, &pubsub.Message{
		Body: data,
		Metadata: map[string]string{
			"topic": topic,
		},
	})
}

func (c *Transport) Listen(ctx context.Context, p *live.PubSub) error {
	sub, err := pubsub.OpenSubscription(ctx, "mem://broadcast")
	if err != nil {
		return fmt.Errorf("could not open subscription: %w", err)
	}

	for {
		msg, err := sub.Receive(ctx)
		if err != nil {
			log.Println("receive message failed: %w", err)
			break
		}

		if c.debug {
			fmt.Println("<< ", string(msg.Body))
		}

		var t live.TransportMessage
		if err := json.Unmarshal(msg.Body, &t); err != nil {
			log.Println("malformed message received: %w", err)
			continue
		}

		if t.Msg.SelfData == nil {
			if t.Msg.Data != nil {
				t.Msg.SelfData = t.Msg.Data
			}
		}

		p.Recieve(t.Topic, t.Msg)
		msg.Ack()
	}
	return fmt.Errorf("stopped receiving messages")
}

func NewPubSub(ctx context.Context, debug bool) (*live.PubSub, error) {
	transport, err := newTransport(ctx, debug)
	if err != nil {
		return nil, err
	}

	return live.NewPubSub(ctx, transport), nil
}
