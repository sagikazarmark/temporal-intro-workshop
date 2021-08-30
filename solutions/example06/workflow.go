package example06

import (
	"math/rand"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Numbers int
}

type Output struct {
	Numbers   []int
	Count     int
	CountOdd  int
	CountEven int
	Sum       int
}

const max = 10

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 06")

	var output Output

	for i := 0; i < input.Numbers; i++ {
		encodedNumber := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return rand.Intn(max)
		})

		var number int

		err := encodedNumber.Get(&number)
		if err != nil {
			return output, err
		}

		workflow.Sleep(ctx, 1*time.Second)

		output.Numbers = append(output.Numbers, number)
		output.Count++
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
