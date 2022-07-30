package main

import (
	"temporal_demo/activity"
	pkgWorkflow "temporal_demo/pkg/workflow"
	"temporal_demo/workflow"
)

func main() {
	cfg := &pkgWorkflow.Config{
		Queue: "abc-test",
	}

	// 初始化worker
	worker, err := pkgWorkflow.NewWorkerClient(cfg)
	if err != nil {
		panic(err)
	}

	// 注册activity
	worker.RegisterActivity(activity.PrintActivity1)
	worker.RegisterActivity(activity.PrintActivity2)

	// 注册workflow
	worker.RegisterWorkflow(workflow.HandleStudent)

	if err := worker.Run(); err != nil {
		panic(err)
	}
}
