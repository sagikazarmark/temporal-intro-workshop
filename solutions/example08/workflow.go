package example08

import (
	"time"

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
	workflow.GetLogger(ctx).Info("starting example 08")

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "workshop",
		ScheduleToCloseTimeout: 3*time.Second + 3*time.Second,
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    3 * time.Second,
		HeartbeatTimeout:       0 * time.Second,
		WaitForCancellation:    false,
		ActivityID:             "",
		RetryPolicy:            nil,
	})

	var activityOutput ActivityOutput

	err := workflow.ExecuteActivity(ctx, Activity08, ActivityInput{input.A, input.B}).Get(ctx, &activityOutput)
	if err != nil {
		return WorkflowOutput{}, err
	}

	return WorkflowOutput{activityOutput.Result}, nil
}
