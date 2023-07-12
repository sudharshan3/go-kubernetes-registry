# Registry-Pod and Image Builder using Kubernetes and Golang
This project demonstrates the use of Kubernetes and Golang to create a registry-pod that stores images, a builder-pod responsible for building and pushing images to the registry-pod via a registry-service, and a pull-pod that pulls the image from the registry-pod. The goal is to showcase a simplified implementation of a container image registry system using Kubernetes and Golang.

## High Level Diagram

![image](https://github.com/sudharshan3/go-kubernetes-registry/assets/50818126/564545a5-0dac-40ea-a5ac-a5838ac541c2)


## Features

- Registry-Pod: A Kubernetes pod that acts as a storage system for container images. It allows you to store and retrieve images within the Kubernetes cluster.
- Builder-Pod: A Kubernetes pod that builds container images using Golang and pushes them to the registry-pod via the registry-service. It demonstrates the process of creating and pushing container images.
- Pull-Pod: A Kubernetes pod that pulls the container image from the registry-pod. It showcases how to retrieve and use container images stored in the registry.

## Prerequisites

- Kubernetes: Ensure you have a running Kubernetes cluster where you can deploy the pods and services.
- Golang: Install Golang on your local machine to build and run the Golang-based components.


## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug fixes, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
