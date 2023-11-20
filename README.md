# Crossplane Top

## Overview

This project provides a Go-based CLI tool to fetch and display metrics (CPU and memory usage) for pods managed by Crossplane in a Kubernetes cluster.
It functions similarly to `kubectl top pods` but is specifically tailored for monitoring Crossplane resources.

## Features

- Fetches real-time CPU and memory metrics for Crossplane-managed pods.
- Filters pods based on Crossplane-specific labels.
- Displays metrics in a user-friendly format.

### Missing Parts

At this point `crossplane-runtime` does not emit any additional metrics to
collect. There is an [open
PR](https://github.com/crossplane/crossplane-runtime/pull/489) adding
`crossplane_resource_drift_seconds` metric and more to follow. `Crossplane-top`
CLI will be able to scrape those additional metrics and add them to the command
output.

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
NAMESPACE            NAME                                                    CPU(cores)   MEMORY(Mi)
crossplane-system    function-go-templating-87197075c1fb-78796c648d-fvkj8    0            6Mi
crossplane-system    provider-helm-59fc94cd939e-cf9f6b97b-rgsk6              1            11Mi
crossplane-system    provider-kubernetes-a1a49ab74384-68ffcc7f45-tpq8k       1            9Mi
-----
```

## Future Customization

You will be able to modify the label selector in the code to target different pods or adjust the namespace as per your Crossplane setup.

<!-- TODO: expose namespace and label selection via CLI -->

## Similar Projects and Inspiration

`Uptest` has an excellent performance testing tooling that creates some custom
performance related metrics, for example [resources readiness](https://github.com/upbound/uptest/blob/6e567ebd9ed30f1b1670d2cbbb679fde9beebc6b/cmd/perf/internal/managed/managed.go#L171).

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your changes.

## License

MIT
