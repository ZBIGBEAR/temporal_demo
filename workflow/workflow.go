package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func HandleStudent(ctx workflow.Context, student interface{}) error {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	result := workflow.ExecuteActivity(ctx, "PrintActivity1", student)

	if err := result.Get(ctx, nil); err != nil {
		return err
	}

	result2 := workflow.ExecuteActivity(ctx, "PrintActivity2", student)

	if err := result2.Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
