package broker

import "github.com/segmentio/kafka-go"

type Producer struct {
	*kafka.Writer
}

func (b *broker) Producer() *Producer {
	return &Producer{
		Writer: &kafka.Writer{
			Addr:        kafka.TCP(b.Addr),
			Balancer:    &kafka.RoundRobin{},
			Compression: kafka.Snappy,
		},
	}
}
