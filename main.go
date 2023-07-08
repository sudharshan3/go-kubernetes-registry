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
	log.Print("Creating Registry Deployment...")
	newRegistry.RegistryName = "registry"
	newRegistry.RegistryServiceName = "registry-service"
	newRegistry.RegistryServicePort = 5000
	newRegistry.RegistryServiceProtocol = "TCP"
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
	log.Print("Creating Registry Service...")
	_, err = newRegistry.CreateRegistryService(kube)
	if err != nil {
		log.Fatal(err)
	}
}
