package workflow

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type WorkerClient interface {
	RegisterActivity(activity interface{})
	RegisterWorkflow(workflow interface{})
	Run() error
}

type workerClient struct {
	w worker.Worker
}

func NewWorkerClient(cfg *Config) (WorkerClient, error) {
	client, err := client.NewLazyClient(client.Options{})
	if err != nil {
		return nil, err
	}

	wc := &workerClient{
		w: worker.New(client, cfg.Queue, worker.Options{}),
	}

	return wc, nil
}

func (w *workerClient) RegisterActivity(activity interface{}) {
	w.w.RegisterActivity(activity)
}

func (w *workerClient) RegisterWorkflow(workflow interface{}) {
	w.w.RegisterWorkflow(workflow)
}

func (w *workerClient) Run() error {
	return w.w.Run(worker.InterruptCh())
}
