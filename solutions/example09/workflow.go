package example09

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkflowInput struct {
	A int
	B int
}

type WorkflowOutput struct {
	Result int
}

func Workflow(ctx workflow.Context, input WorkflowInput) (WorkflowOutput, error) {
	workflow.GetLogger(ctx).Info("starting example 09")

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "workshop",
		ScheduleToCloseTimeout: 3*time.Second + 3*time.Second,
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    3 * time.Second,
		HeartbeatTimeout:       0 * time.Second,
		WaitForCancellation:    false,
		ActivityID:             "",
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    5,
		},
	})

	var activityOutput ActivityOutput

	err := workflow.ExecuteActivity(ctx, Activity09, ActivityInput{input.A, input.B}).Get(ctx, &activityOutput)
	if err != nil {
		return WorkflowOutput{}, err
	}

	return WorkflowOutput{activityOutput.Result}, nil
}
