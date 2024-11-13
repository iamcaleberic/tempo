package main

import (
	"github.com/iamcaleberic/tempo/activity"
	loglib "github.com/iamcaleberic/tempo/logger"
	"github.com/iamcaleberic/tempo/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

var (
	logger = loglib.InitLogger()
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		Logger: loglib.NewZapAdapter(logger),
	})
	if err != nil {
		logger.Fatal("unable to create Temporal client", zap.Error(err))
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	k8sWorker := worker.New(c, workflow.K8sTaskQueue, worker.Options{})
	k8sWorker.RegisterWorkflow(workflow.ListPodsWorkflow)
	k8sWorker.RegisterActivity(activity.ListPods)

	// start k8s worker
	// go func() {
	err = k8sWorker.Run(worker.InterruptCh())
	if err != nil {
		logger.Fatal("unable to start Worker", zap.Error(err))
	}
	// }()
}
