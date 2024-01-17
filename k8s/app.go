package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/nemcikjan/app-balancer/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
// kubeConfigFilename = "kube-config.yml"
// available_colors   = [3]string{"red", "green", "blue"}
)

type K8sAppClient struct {
	clientset *kubernetes.Clientset
}

func NewK8sAppClient() (*K8sAppClient, error) {
	kubeConfigFilename, found := os.LookupEnv("KUBE_CONFIG_FILE")

	var kubeConfig *rest.Config

	if found {
		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		kubeConfig, err = clientcmd.BuildConfigFromFlags("", currentPath+"/"+kubeConfigFilename)
		if err != nil {
			panic(err.Error())
		}
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error creating in-cluster config: %v\n", err)
			os.Exit(1)
		}
		kubeConfig = config
	}
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	return &K8sAppClient{
		clientset: clientset,
	}, nil
}

func (c K8sAppClient) CreateApp(id string) {
	pod := createPod(id)
	_, err := c.clientset.CoreV1().Pods("tasks").Create(context.Background(), pod, metav1.CreateOptions{})
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

// func createJob(id string) *batchv1.Job {
// 	resourceRequirements := corev1.ResourceRequirements{
// 		Requests: corev1.ResourceList{
// 			corev1.ResourceCPU:    resource.MustParse(fmt.Sprintf("%dm", utils.RandomNumber(5, 500))),    // Request 500 milliCPUs
// 			corev1.ResourceMemory: resource.MustParse(fmt.Sprintf("%dMi", utils.RandomNumber(10, 1000))), // Request 1 GiB of memory
// 		},
// 	}
// 	job := &batchv1.Job{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: fmt.Sprintf("pod-%s", id),
// 			// Annotations: map[string]string{"v2x.context/speed": ""},
// 			Labels: map[string]string{"v2x": "true"},
// 		},
// 		Spec: batchv1.JobSpec{
// 			Template: corev1.PodTemplateSpec{
// 				Spec: corev1.PodSpec{
// 					RestartPolicy: corev1.RestartPolicyNever,
// 					Containers: []corev1.Container{
// 						{
// 							Name:      "task",
// 							Image:     "alpine@sha256:6ce9a9a256a3495ae60ab0059ed1c7aee5ee89450477f2223f6ea7f6296df555",
// 							Command:   []string{"/bin/sh"},
// 							Args:      []string{"-c", fmt.Sprintf("sleep %d && exit 0", utils.RandomNumber(5, 100))},
// 							Resources: resourceRequirements,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	return job
// }

func createPod(id string) *corev1.Pod {
	resourceRequirements := corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(fmt.Sprintf("%dm", utils.RandomNumber(5, 500))),    // Request 500 milliCPUs
			corev1.ResourceMemory: resource.MustParse(fmt.Sprintf("%dMi", utils.RandomNumber(10, 1000))), // Request 1 GiB of memory
		},
	}
	execTime := utils.RandomNumber(2, 100)
	var color string
	if execTime <= 10 {
		color = "red"
	} else if execTime <= 50 {
		color = "green"
	} else {
		color = "blue"
	}
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("task-%s", id),
			Annotations: map[string]string{
				"v2x.context/priority":  fmt.Sprintf("%d", utils.RandomNumber(1, 5)),
				"v2x.context/color":     color,
				"v2x.context/exec_time": fmt.Sprintf("%d", execTime),
			},
			Labels: map[string]string{"v2x": "true"},
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:      "task",
					Image:     "alpine@sha256:6ce9a9a256a3495ae60ab0059ed1c7aee5ee89450477f2223f6ea7f6296df555",
					Command:   []string{"/bin/sh"},
					Args:      []string{"-c", fmt.Sprintf("sleep %d && exit 0", execTime)},
					Resources: resourceRequirements,
				},
			},
		},
	}
	return pod
}
