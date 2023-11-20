# Crossplane Top

## Overview

This project provides a Go-based CLI tool to fetch and display metrics (CPU and memory usage) for pods managed by Crossplane in a Kubernetes cluster.
It functions similarly to `kubectl top pods` but is specifically tailored for monitoring Crossplane resources.

## Features

- Fetches real-time CPU and memory metrics for Crossplane-managed pods.
- Filters pods based on Crossplane-specific labels.
- Displays metrics in a user-friendly format.

## Prerequisites

- Go (version 1.15 or higher)
- Access to a Kubernetes cluster with Crossplane installed.
- Kubernetes config file set up (typically located at `~/.kube/config`).

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/Piotr1215/crossplane-top
   ```
2. Navigate to the cloned directory:
   ```
   cd your-repo
   ```
3. Build the project (optional):
   ```
   go build -o crossplane-top
   ```

## Usage

Run the program directly using Go:

```
go run main.go
```

Or, if you have built the binary:

```
./crossplane-top
```

The tool will output the CPU and memory usage of each Crossplane-managed pod in your Kubernetes cluster.
This includes `providers` and `functions`

## Example Output

```
Pod Name: provider-helm-59fc94cd939e-6587d548cf-vgcxc
Container Name: provider-helm
CPU Usage: 1m
Memory Usage: 12Mi
-----
Pod Name: provider-kubernetes-a1a49ab74384-6dd6c6fdf8-w2bpv
Container Name: provider-kubernetes
CPU Usage: 1m
Memory Usage: 9Mi
-----
```

## Customization

You can modify the label selector in the code to target different pods or adjust the namespace as per your Crossplane setup.

<!-- TODO: expose namespace and label selection via CLI -->

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your changes.

## License

MIT
