package example10

import (
	"context"
	"errors"

	"go.temporal.io/sdk/activity"
)

type ActivityInput struct {
	A int
	B int
}

type ActivityOutput struct {
	Result int
}

func Activity10(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
	activityInfo := activity.GetInfo(ctx)

	if activityInfo.Attempt < 3 {
		return ActivityOutput{}, errors.New("attempts under 3")
	}

	return ActivityOutput{input.A + input.B}, nil
}
