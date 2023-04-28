package asyncQueue

import (
	"github.com/hibiken/asynq"
)

type AsyncQueue interface {
	DistributeTask(taskName string, value []byte) error
	ProcessTask(taskName string, handler asynq.HandlerFunc) error
	Close() error
}

type asyncQueue struct {
	Distributor asynq.Client
	Processor   asynq.Server
}

func New(address string, password string) *asyncQueue {
	opts := asynq.RedisClientOpt{
		Addr:     address,
		Password: password,
	}
	cfg := asynq.Config{
		Concurrency: 10,
	}
	return &asyncQueue{
		Distributor: *asynq.NewClient(opts),
		Processor:   *asynq.NewServer(opts, cfg),
	}
}

func (a *asyncQueue) Close() error {
	a.Processor.Stop()
	return a.Distributor.Close()
}

var _ AsyncQueue = (*asyncQueue)(nil)
