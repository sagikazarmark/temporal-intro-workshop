package example01

import "go.temporal.io/sdk/workflow"

func Workflow(ctx workflow.Context, a int, b int) (int, error) {
	return a + b, nil
}
