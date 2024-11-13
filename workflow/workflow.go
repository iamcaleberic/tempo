package workflow

import (
	"time"

	"github.com/iamcaleberic/tempo/activity"
	"go.temporal.io/sdk/workflow"
)

func ListPodsWorkflow(ctx workflow.Context) ([]string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	logger := workflow.GetLogger(ctx)
	logger.Info("ListPodsWorkflow started")

	ctx = workflow.WithActivityOptions(ctx, options)

	var result []string
	err := workflow.ExecuteActivity(ctx, activity.ListPods).Get(ctx, &result)

	return result, err
}
