package example05

import (
	"math/rand"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Number int
}

type Output struct {
	Result int
}

const max = 10

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 05")

	number := input.Number

	if number < 1 {
		workflow.GetLogger(ctx).Info("generating random number")

		encodedNumber := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return rand.Intn(max)
		})

		err := encodedNumber.Get(&number)
		if err != nil {
			return Output{}, err
		}
	}

	result := 1

	for i := 1; i <= number; i++ {
		workflow.Sleep(ctx, 5*time.Second)

		result *= i
	}

	output := Output{
		Result: result,
	}

	return output, nil
}
