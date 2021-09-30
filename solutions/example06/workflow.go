package example06

import (
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 06")

	var number int

	err := workflow.SetQueryHandler(ctx, "current_number", func() (int, error) {
		return number, nil
	})
	if err != nil {
		return err
	}

	signalChan := workflow.GetSignalChannel(ctx, "set_number")

	s := workflow.NewSelector(ctx)

	s.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &number)

		workflow.GetLogger(ctx).Info("Received number", "number", number)
	})

	for {
		s.Select(ctx)
	}
}
