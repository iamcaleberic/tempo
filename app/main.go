package main

import (
	"context"

	loglib "github.com/iamcaleberic/tempo/logger"
	"github.com/iamcaleberic/tempo/workflow"
	"go.uber.org/zap"

	"go.temporal.io/sdk/client"
)

var (
	logger = loglib.InitLogger()
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		logger.Fatal("unable to create Temporal client", zap.Error(err))
	}
	defer c.Close()

	listPodsWorklfowOptions := client.StartWorkflowOptions{
		ID:        "listpods-workflow",
		TaskQueue: workflow.K8sTaskQueue,
	}

	// list-pods Workflow
	lw, err := c.ExecuteWorkflow(context.Background(), listPodsWorklfowOptions, workflow.ListPodsWorkflow)
	if err != nil {
		logger.Fatal("unable to complete Workflow", zap.Error(err))
	}

	// Get the results
	var listPodsResult []string
	err = lw.Get(context.Background(), &listPodsResult)
	if err != nil {
		logger.Fatal("unable to get Workflow result", zap.Error(err))
	}

	logger.Info("listpod-result", zap.String("WorkflowID", lw.GetID()), zap.String("runID", lw.GetRunID()))
	logger.Info("pods", zap.Strings("", listPodsResult))

}
