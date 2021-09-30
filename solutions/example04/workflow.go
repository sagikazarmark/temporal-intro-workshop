package example04

import (
	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Number int
}

type Output struct {
	Result int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 04")

	if input.Number < 1 {
		workflow.GetLogger(ctx).Info("invalid number", "number", input.Number)
	}

	result := 1

	for i := 1; i <= input.Number; i++ {
		result *= i
	}

	output := Output{
		Result: result,
	}

	return output, nil
}
