package workflow

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"
)

type Config struct {
	ID    string
	Queue string
}

type WorkflowClient interface {
	ExecuteWorkflow(ctx context.Context, handler interface{}, args ...interface{}) error
	StartCron(ctx context.Context, schedule string, handler interface{}, args ...interface{}) error
}

type workflowClient struct {
	c                    client.Client
	startWorkflowOptions *client.StartWorkflowOptions
}

func NewWorkflowClient(cfg *Config) (WorkflowClient, error) {
	wfc := &workflowClient{}
	var err error
	wfc.c, err = client.NewLazyClient(client.Options{})
	if err != nil {
		return nil, err
	}

	wfc.startWorkflowOptions = &client.StartWorkflowOptions{
		ID:        cfg.ID,
		TaskQueue: cfg.Queue,
	}

	return wfc, nil
}

func (c *workflowClient) ExecuteWorkflow(ctx context.Context, handler interface{}, args ...interface{}) error {
	// 开始执行
	result, err := c.c.ExecuteWorkflow(ctx, *c.startWorkflowOptions, handler, args...)
	if err != nil {
		return err
	}

	// Get方法会阻塞
	fmt.Printf("ID:%s, RunID:%s", result.GetID(), result.GetRunID())
	return nil
}

func (c *workflowClient) StartCron(ctx context.Context, schedule string, handler interface{}, args ...interface{}) error {
	// 开始执行
	option := client.StartWorkflowOptions{
		ID:           c.startWorkflowOptions.ID,
		TaskQueue:    c.startWorkflowOptions.TaskQueue,
		CronSchedule: schedule,
	}
	result, err := c.c.ExecuteWorkflow(ctx, option, handler, args...)
	if err != nil {
		return err
	}

	// Get方法会阻塞
	fmt.Printf("ID:%s, RunID:%s", result.GetID(), result.GetRunID())
	return nil
}
