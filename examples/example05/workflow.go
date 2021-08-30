package example05

import (
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 05")

	return nil
}
