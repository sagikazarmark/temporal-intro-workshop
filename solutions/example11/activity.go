package example11

import (
	"context"
	"errors"

	"go.temporal.io/sdk/activity"
)

type ActivityInput struct {
	Number int
}

type ActivityOutput struct {
	Result int
}

func Activity11(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
	activityInfo := activity.GetInfo(ctx)

	if activityInfo.Attempt < 1 {
		return ActivityOutput{}, errors.New("first attempt")
	}

	result := 1

	for i := 1; i <= input.Number; i++ {
		result *= i
	}

	return ActivityOutput{result}, nil
}
