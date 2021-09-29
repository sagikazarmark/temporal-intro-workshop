package main

import (
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example01"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example03"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example04"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example05"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example06"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example07"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example08"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example09"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example10"
	"github.com/sagikazarmark/temporal-intro-workshop/examples/example11"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func register(w worker.Registry) {
	w.RegisterWorkflow(example01.Workflow)
	w.RegisterWorkflowWithOptions(example01.Workflow, workflow.RegisterOptions{Name: "example01"})

	w.RegisterWorkflowWithOptions(example03.Workflow, workflow.RegisterOptions{Name: "example03"})

	w.RegisterWorkflowWithOptions(example04.Workflow, workflow.RegisterOptions{Name: "example04"})

	w.RegisterWorkflowWithOptions(example05.Workflow, workflow.RegisterOptions{Name: "example05"})

	w.RegisterWorkflowWithOptions(example06.Workflow, workflow.RegisterOptions{Name: "example06"})

	w.RegisterWorkflowWithOptions(example07.Workflow, workflow.RegisterOptions{Name: "example07"})

	w.RegisterWorkflowWithOptions(example08.Workflow, workflow.RegisterOptions{Name: "example08"})

	w.RegisterWorkflowWithOptions(example09.Workflow, workflow.RegisterOptions{Name: "example09"})
	w.RegisterActivity(example09.Activity09)
	w.RegisterActivityWithOptions(example09.Activity09, activity.RegisterOptions{Name: "example09"})

	w.RegisterWorkflowWithOptions(example10.Workflow, workflow.RegisterOptions{Name: "example10"})
	w.RegisterActivity(example10.Activity10)
	w.RegisterActivityWithOptions(example10.Activity10, activity.RegisterOptions{Name: "example10"})

	w.RegisterWorkflowWithOptions(example11.Workflow, workflow.RegisterOptions{Name: "example11"})
	w.RegisterActivity(example11.Activity11)
	w.RegisterActivityWithOptions(example11.Activity11, activity.RegisterOptions{Name: "example11"})
}
