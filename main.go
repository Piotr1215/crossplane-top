package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset for Kubernetes
	k8sClientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset for Metrics
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Fetch all pods in the Crossplane namespace
	pods, err := k8sClientset.CoreV1().Pods("crossplane-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Filter pods based on label key prefix
	for _, pod := range pods.Items {
		for labelKey := range pod.GetLabels() {
			if strings.HasPrefix(labelKey, "pkg.crossplane.io/provider") {
				// Get metrics for this pod
				podMetrics, err := metricsClientset.MetricsV1beta1().PodMetricses(pod.Namespace).Get(context.TODO(), pod.Name, metav1.GetOptions{})
				if err != nil {
					fmt.Printf("Error getting metrics for pod %s: %v\n", pod.Name, err)
					continue
				}

				// Print pod metrics
				fmt.Printf("Pod Name: %s\n", podMetrics.Name)
				for _, container := range podMetrics.Containers {
					fmt.Printf("Container Name: %s\n", container.Name)
					fmt.Printf("CPU Usage: %dm\n", container.Usage.Cpu().ScaledValue(resource.Milli))
					fmt.Printf("Memory Usage: %dMi\n", container.Usage.Memory().ScaledValue(resource.Mega))
				}
				fmt.Println("-----")
			}
		}
	}
}

// homeDir returns the home directory for the current user
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
