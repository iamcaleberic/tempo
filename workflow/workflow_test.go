package workflow

import (
	"testing"

	"github.com/iamcaleberic/tempo/activity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_ListPodsWorkflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(activity.ListPods, mock.Anything).Return([]string{"test-pod-1", "test-pod-2"}, nil)

	env.ExecuteWorkflow(ListPodsWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var podNames []string
	require.NoError(t, env.GetWorkflowResult(&podNames))
	require.Equal(t, []string{"test-pod-1", "test-pod-2"}, podNames)
}
