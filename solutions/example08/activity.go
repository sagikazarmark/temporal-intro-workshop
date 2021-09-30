package example08

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

func Activity08(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
	return ActivityOutput{input.A + input.B}, nil
}
