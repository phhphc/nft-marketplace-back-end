package asyncQueue

import (
	"time"

	"github.com/hibiken/asynq"
)

func (aq *asyncQueue) DistributeTask(taskName string, value []byte) error{
	task := asynq.NewTask(taskName, value)
	// Task will be processed up to 6 times, 20 seconds break between 2 times
	_, err := aq.Distributor.Enqueue(task, asynq.MaxRetry(6), asynq.Timeout(20 * time.Second))

	if err != nil {
		return err
	}
	return nil	
}

func (aq *asyncQueue) ProcessTask(taskName string, handler asynq.HandlerFunc) error{
	mux := asynq.NewServeMux()
	mux.HandleFunc(taskName, handler)	
	err := aq.Processor.Run(mux)

	if err != nil {
		return err
	}
	return nil
}
