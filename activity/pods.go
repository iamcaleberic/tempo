package activity

import (
	"context"

	"github.com/iamcaleberic/tempo/k8s"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListPodsObject struct{}

// ListPods lists all pods in current selected context
func ListPods(ctx context.Context, params ListPodsObject) ([]string, error) {
	logger.Info("Activity started.")
	var podNames []string

	clientSet, err := k8s.CreateClientSet(ctx)
	if err != nil {
		logger.Error("failed to create clientSet", zap.Error(err))
		return podNames, err
	}

	pods, err := clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		logger.Error("failed to list pods", zap.Error(err))
		return podNames, err
	}

	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames, nil
}
