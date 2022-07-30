package main

import (
	"context"
	"temporal_demo/entity"
	pkgWorkflow "temporal_demo/pkg/workflow"
)

func main() {
	student1 := &entity.Student{
		Name: "jobs",
		Age:  18,
	}

	cfg := &pkgWorkflow.Config{
		Queue: "abc-test",
	}

	wf, err := pkgWorkflow.NewWorkflowClient(cfg)
	if err != nil {
		panic(err)
	}
	// 执行
	if err := wf.ExecuteWorkflow(context.Background(), "HandleStudent", student1); err != nil {
		panic(err)
	}

}
