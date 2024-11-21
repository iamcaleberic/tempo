package k8s

import (
	"context"
	"flag"
	"path/filepath"

	loglib "github.com/iamcaleberic/tempo/logger"
	"go.uber.org/zap"

	"k8s.io/client-go/kubernetes"
	// load auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	logger = loglib.InitLogger()
)

func CreateClientSet(ctx context.Context) (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.Error("failed to build client cmd", zap.Error(err))
		return &kubernetes.Clientset{}, err
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error("failed to create clientSet", zap.Error(err))
		return &kubernetes.Clientset{}, err
	}

	return clientSet, nil

}
