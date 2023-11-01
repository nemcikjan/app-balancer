package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/nemcikjan/app-balancer/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfigFilename = "kube-config.yml"
)

type K8sAppClient struct {
	clientset *kubernetes.Clientset
}

func NewK8sAppClient() (*K8sAppClient, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	config, err := clientcmd.BuildConfigFromFlags("", currentPath+"/"+kubeConfigFilename)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &K8sAppClient{
		clientset: clientset,
	}, nil
}

func (c K8sAppClient) CreateApp(id string) {
	pod := createPod(id)
	_, err := c.clientset.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating pod: %v\n", err)
		// panic(err.Error())
	}
}

func (c K8sAppClient) UpdateApp(id string, speed float64) {
	// Fetch the existing pod
	pod, err := c.clientset.CoreV1().Pods("default").Get(context.TODO(), fmt.Sprintf("pod-%s", id), metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error fetching pod: %v\n", err)
		// os.Exit(1)
	}
	if pod.Status.Phase != corev1.PodSucceeded {
		pod.Annotations["v2x.context/speed"] = fmt.Sprintf("%f", speed)
		// Update the pod with the new annotation
		_, err = c.clientset.CoreV1().Pods("default").Update(context.TODO(), pod, metav1.UpdateOptions{
			FieldManager: "v2x",
		})
		if err != nil {
			fmt.Printf("Error updating pod: %v\n", err)
			// os.Exit(1)
		}

		fmt.Println("Pod annotation updated successfully!")
	} else {
		fmt.Printf("Skipping update, pod %s not running\n", id)
	}

}

func createPod(id string) *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprintf("pod-%s", id),
			Annotations: map[string]string{"v2x.context/speed": ""},
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:    "task",
					Image:   "alpine@sha256:6ce9a9a256a3495ae60ab0059ed1c7aee5ee89450477f2223f6ea7f6296df555",
					Command: []string{"/bin/sh", "-c"},
					Args:    []string{fmt.Sprintf("sleep %d && exit 0", utils.RandomNumber(100))},
				},
			},
		},
	}
	return pod
}
