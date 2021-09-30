package example07

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Number int
}

type Output struct {
	Result int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 07")

	result := 1

	err := workflow.SetQueryHandler(ctx, "current_result", func() (int, error) {
		return result, nil
	})
	if err != nil {
		return Output{}, err
	}

	for i := 1; i <= input.Number; i++ {
		workflow.Sleep(ctx, 10*time.Second)

		result *= i
	}

	output := Output{
		Result: result,
	}

	return output, nil
}
