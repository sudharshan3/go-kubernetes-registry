package main

import (
	"log"

	"github.com/sudharshan3/go-kubernetes-registry/kubeconnector"
)

func main() {
	log.Print("Connecting to Kubernetes Cluster...")
	kube, err := kubeconnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection Established Successfully...")
	log.Print(kube)
}
