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

func getKubeConfig() (string, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		return kubeconfig, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}
	return filepath.Join(home, ".kube", "config"), nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {

	// Build the config from the kubeconfig path
	kubeconfig, err := getKubeConfig()
	if err != nil {
		return fmt.Errorf("failed to get kubeconfig: %w", err)
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("could not build config from flags: %w", err)
	}

	// Create the clientset for Kubernetes
	k8sClientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("could not create the clientset for Kubernetes: %w", err)
	}

	// Create the clientset for Metrics
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("could not create the clientset for Metrics: %w", err)
	}

	// Fetch all pods from all namespaces in case of Crossplane pods being installed elswhere
	pods, err := k8sClientset.CoreV1().Pods(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("could not fetch all pods from all namespaces: %w", err)
	}

	fmt.Printf("%-20s %-55s %-12s %-15s\n", "NAMESPACE", "NAME", "CPU(cores)", "MEMORY(Mi)")

	// Loop through pods and print metrics
	for _, pod := range pods.Items {
		for labelKey := range pod.GetLabels() {
			if strings.HasPrefix(labelKey, "pkg.crossplane.io/provider") || strings.HasPrefix(labelKey, "pkg.crossplane.io/function") {
				podMetrics, err := metricsClientset.MetricsV1beta1().PodMetricses(pod.Namespace).Get(context.TODO(), pod.Name, metav1.GetOptions{})
				if err != nil {
					fmt.Printf("Error getting metrics for pod %s: %v\n", pod.Name, err)
					continue
				}

				for _, container := range podMetrics.Containers {
					cpuUsage := container.Usage.Cpu().ScaledValue(resource.Milli)
					memoryUsage := fmt.Sprintf("%dMi", container.Usage.Memory().ScaledValue(resource.Mega))
					fmt.Printf("%-20s %-55s %-12d %-15s\n", pod.Namespace, pod.Name, cpuUsage, memoryUsage)
				}
			}
		}
	}
	return nil
}
