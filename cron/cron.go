package cron

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func CronJob(ctx workflow.Context, args map[string]interface{}) error {
	fmt.Printf("[CronJob] args:%v", args)
	return nil
}
