package registry

import (
	"context"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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

func (registryDetails *Registry) CreateRegistryService(clientSet *kubernetes.Clientset) (*corev1.Service, error) {
	if registryDetails.RegistryServiceName == "" {
		log.Fatal("Please Provide Registry Service Name")
	}
	if registryDetails.RegistryServicePort == 0 {
		log.Fatal("Please Provide Registry Service Port")
	}
	if registryDetails.RegistryServiceProtocol == "" {
		log.Fatal("Please Provide Registry Service Protocol")
	}
	registryService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      registryDetails.RegistryServiceName,
			Namespace: registryDetails.RegistryNamespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Selector: map[string]string{
				"app": registryDetails.RegistryAppName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       int32(registryDetails.RegistryServicePort),
					TargetPort: intstr.IntOrString{IntVal: int32(registryDetails.RegistryServicePort)},
					Protocol:   corev1.Protocol(registryDetails.RegistryServiceProtocol),
				},
			},
		},
	}
	service, err := clientSet.CoreV1().Services(registryDetails.RegistryNamespace).Create(context.TODO(), registryService, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return service, nil
}
