package example09

import (
	"context"
)

type ActivityInput struct {
	A int
	B int
}

type ActivityOutput struct {
	Result int
}

func Activity09(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
	return ActivityOutput{input.A + input.B}, nil
}
