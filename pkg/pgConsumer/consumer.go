package pgConsumer

import "github.com/segmentio/kafka-go"

type Consumer struct {
	*kafka.Reader
}

func (b *broker) Consumer(topic string, groupId string) *Consumer {
	return &Consumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{b.Addr},
			GroupID:  groupId,
			Topic:    topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
	}
}
