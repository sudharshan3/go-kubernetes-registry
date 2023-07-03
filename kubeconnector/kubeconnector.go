package kubeconnector

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Connect() (*kubernetes.Clientset, error) {
	log.Println("Get Home Directory...")
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	log.Println("Get Kubernetes Config...")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	log.Println("Creating Clientset...")
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return clientSet, nil
}
