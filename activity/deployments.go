package activity

import "context"

type GetDeploymentObject struct {
	Name      string
	Namespace string
}

func GetDeployment(ctx context.Context, params GetDeploymentObject) {
	logger.Info("Activity started.")
}
