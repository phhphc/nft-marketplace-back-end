package broker

type Broker interface {
	Consumer(topic string, groupId string) *Consumer
	Producer() *Producer
}

type broker struct {
	Addr string
}

func New(addr string) *broker {
	return &broker{Addr: addr}
}

var _ Broker = (*broker)(nil)
