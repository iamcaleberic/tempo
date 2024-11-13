package workflow

import (
	loglib "github.com/iamcaleberic/tempo/logger"
)

var (
	logger = loglib.InitLogger()
)

const K8sTaskQueue = "K8S_TASK_QUEUE"
