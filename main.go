package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset for Metrics
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Fetch node metrics
	nodeMetricsList, err := metricsClientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Print details
	for _, nodeMetric := range nodeMetricsList.Items {
		fmt.Printf("Node Name: %s\n", nodeMetric.Name)
		fmt.Printf("CPU Usage: %s\n", nodeMetric.Usage.Cpu().String())
		fmt.Printf("Memory Usage: %s\n", nodeMetric.Usage.Memory().String())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
