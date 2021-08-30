package example05

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *WorkflowTestSuite) Test_Success() {
	s.env.RegisterWorkflow(Workflow)
	s.env.ExecuteWorkflow(Workflow, Input{[]int{0, 1, 2, 3, 4, 5}})

	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())

	var output Output
	s.Require().NoError(s.env.GetWorkflowResult(&output))

	expectedOutput := Output{
		Count:     6,
		CountOdd:  3,
		CountEven: 3,
		Sum:       15,
	}

	s.Equal(expectedOutput, output)
}

func (s *WorkflowTestSuite) Test_NoNumbers() {
	s.env.RegisterWorkflow(Workflow)
	s.env.ExecuteWorkflow(Workflow, Input{})

	s.Require().True(s.env.IsWorkflowCompleted())

	var appError *temporal.ApplicationError
	if err := s.env.GetWorkflowError(); !errors.As(err, &appError) {
		s.Failf("unexpected error", "unexpected error: %#v", err)
	}

	s.EqualError(appError, "no numbers")
}
