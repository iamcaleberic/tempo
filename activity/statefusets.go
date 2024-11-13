package activity

import (
	"context"
)

type GetStatefulSetObject struct {
	Name      string
	Namespace string
}

func GetStatefulSet(ctx context.Context, params GetStatefulSetObject) {
	logger.Info("Activity started.")

}
