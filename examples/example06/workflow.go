package example06

import (
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 06")

	return nil
}
