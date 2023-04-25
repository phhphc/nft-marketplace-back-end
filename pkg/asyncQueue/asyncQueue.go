package asyncQueue

import "github.com/hibiken/asynq"


type AsyncQueue interface {
	DistributeTask(taskName string, value []byte) error
	ProcessTask(taskName string, handler asynq.HandlerFunc) error
}

type asyncQueue struct {
	Distributor asynq.Client
 	Processor 	asynq.Server
}

func New(address string, password string) *asyncQueue {
	opts := asynq.RedisClientOpt{
		Addr: address,
		Password: password,
	}
	cfg := asynq.Config{Concurrency: 10}
	return &asyncQueue{
		Distributor: *asynq.NewClient(opts),
		Processor: *asynq.NewServer(opts, cfg),
	}
}

var _ AsyncQueue = (*asyncQueue)(nil)