package example11

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/activity"
)

type ActivityInput struct {
	Number int
}

type ActivityOutput struct {
	Even bool
}

func Activity11(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
	activityInfo := activity.GetInfo(ctx)

	if activityInfo.Attempt < 1 {
		return ActivityOutput{}, errors.New("first attempt")
	}

	if activityInfo.Attempt < 2 {
		time.Sleep(5 * time.Second)
	}

	return ActivityOutput{input.Number%2 == 0}, nil
}
