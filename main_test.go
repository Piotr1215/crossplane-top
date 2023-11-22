package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetKubeConfig(t *testing.T) {
	// Test when KUBECONFIG environment variable is set
	os.Setenv("KUBECONFIG", "/path/to/kubeconfig")
	expected := "/path/to/kubeconfig"
	actual, err := getKubeConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}

	// Test when KUBECONFIG environment variable is not set
	os.Unsetenv("KUBECONFIG")
	homeDir, _ := os.UserHomeDir()
	expected = filepath.Join(homeDir, ".kube", "config")
	actual, err = getKubeConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}
