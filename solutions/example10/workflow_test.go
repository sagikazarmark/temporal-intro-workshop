package example10

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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
	s.env.RegisterActivity(Activity10)

	s.env.OnActivity(Activity10, mock.Anything, ActivityInput{5}).Return(ActivityOutput{120}, nil)

	s.env.ExecuteWorkflow(Workflow, WorkflowInput{5})

	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())

	var output WorkflowOutput
	s.Require().NoError(s.env.GetWorkflowResult(&output))

	expectedOutput := WorkflowOutput{
		Result: 120,
	}

	s.Equal(expectedOutput, output)
}
