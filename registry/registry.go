package registry

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Registry struct {
	RegistryName            string
	RegistryNamespace       string
	RegistryImage           string
	RegistryPort            int32
	RegistryAppName         string
	RegistryMountPath       string
	RegistryReplicas        *int32
	RegistryServiceName     string
	RegistryServicePort     int
	RegistryServiceProtocol string
}

func (registryDetails *Registry) CreateRegistryDeployment(clientSet *kubernetes.Clientset) (*appsv1.Deployment, error) {
	registryDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      registryDetails.RegistryName,
			Namespace: registryDetails.RegistryNamespace,
			Labels: map[string]string{
				"app": registryDetails.RegistryAppName,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: registryDetails.RegistryReplicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": registryDetails.RegistryAppName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": registryDetails.RegistryAppName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  registryDetails.RegistryName,
							Image: registryDetails.RegistryImage,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: registryDetails.RegistryPort,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "registry-data",
									MountPath: registryDetails.RegistryMountPath,
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "registry-data",
						},
					},
				},
			},
		},
	}
	deployment, err := clientSet.AppsV1().Deployments(registryDetails.RegistryNamespace).Create(context.TODO(), registryDeployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
