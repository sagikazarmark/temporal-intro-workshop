package example08

import (
	"errors"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Numbers []int
}

type Output struct {
	Count     int
	CountOdd  int
	CountEven int
	Sum       int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 08")

	output := Output{
		Count: len(input.Numbers),
	}

	if output.Count == 0 {
		return output, errors.New("no numbers")
	}

	var currentNumber int

	err := workflow.SetQueryHandler(ctx, "current_number", func() (int, error) {
		return currentNumber, nil
	})
	if err != nil {
		return output, err
	}

	for _, number := range input.Numbers {
		currentNumber = number

		workflow.Sleep(ctx, 10*time.Second)

		output.Sum += number

		if number%2 == 0 {
			output.CountEven++
		} else {
			output.CountOdd++
		}

		if number%3 == 0 {
			workflow.GetLogger(ctx).Info("number divisible by 3", "number", number)
		}
	}

	return output, nil
}
