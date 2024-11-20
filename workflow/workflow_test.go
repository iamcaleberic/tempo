package workflow

import (
	"testing"

	"github.com/iamcaleberic/tempo/activity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func Test_ListPodsWorkflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(activity.ListPods, mock.Anything, activity.ListPodsObject{}).Return([]string{"test-pod-1", "test-pod-2"}, nil)

	env.ExecuteWorkflow(ListPodsWorkflow)
	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())

	var podNames []string
	assert.NoError(t, env.GetWorkflowResult(&podNames))
	assert.Equal(t, []string{"test-pod-1", "test-pod-2"}, podNames)
}
