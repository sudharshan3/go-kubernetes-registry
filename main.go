package main

import (
	"log"

	"github.com/sudharshan3/go-kubernetes-registry/kubeconnector"
	"github.com/sudharshan3/go-kubernetes-registry/registry"
)

func main() {
	var newRegistry registry.Registry
	var replicas int32 = 1
	log.Print("Connecting to Kubernetes Cluster...")
	kube, err := kubeconnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection Established Successfully...")
	log.Print("Create Registry Deployment...")
	newRegistry.RegistryName = "registry"
	newRegistry.RegistryAppName = "registry"
	newRegistry.RegistryNamespace = "default"
	newRegistry.RegistryImage = "registry:2"
	newRegistry.RegistryPort = 5000
	newRegistry.RegistryReplicas = &replicas
	newRegistry.RegistryMountPath = "/var/lib/registry"
	_, err = newRegistry.CreateRegistryDeployment(kube)
	if err != nil {
		log.Fatal(err)
	}

}
