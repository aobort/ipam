# k8s-subnet-machine-request
k8s operator for SubnetMachineRequest CRD

### CRD parameters

| Parameter  | Description | Example | Validation rules |
| ------------- | ------------- | ------------- | ------------- |
| subnet | Subnet reference | subnet | Should exist |
| machineRequest | Machine Request reference | machinerequest1 | Should exist |
| ip | IP to request | 10.12.34.64 | IPv4 or IPv6 - should be available in specified subnet |

## Getting started

### Required tools

Following tools are required to make changes on that package.

- k8s cluster access to deploy and test the result (via minikube or docker desktop locally)
- [make](https://www.gnu.org/software/make/) - to execute build goals
- [golang](https://golang.org/) - to compile the code
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - to interact with k8s cluster via CLI
- [kustomize](https://kustomize.io/) - to generate deployment configs
- [kubebuilder](https://book.kubebuilder.io) - framework to build operators
- [operator framework](https://operatorframework.io/) - framework to maintain project structure
- [helm](https://helm.sh/) - to work with helm charts

### Install definitions

In order to build and deploy, execute following command set: `make install`

### Development

This repo references other CRDs and you need to install them to proceed:
- Subnet https://github.com/onmetal/k8s-subnet
- Machine Requests https://github.com/onmetal/k8s-machine-requests

So you might need to run `go env -w GOPRIVATE='github.com/onmetal/*'` first to build it.

So to run controller for development without deploy do: `make run`

### Deploy 

Docker registry is required to build and push an image. 
For local development you can use local registry e.g. `localhost:5000` for [docker desktop](https://docs.docker.com/registry/deploying/).

Replace with your registry if you're using quay or anything else.

```
# ! Be sure to install CRDs first
# Build and push Docker image
make docker-build docker-push IMG="localhost:5000/k8s-subnet-machine-request:latest" GIT_USER=yourusername GIT_PASSWORD=youraccesstoken

# Deploy
make deploy IMG="localhost:5000/k8s-subnet-machine-request:latest"
```

### Helm chart

TODO

### Use

`./config/samples/` directory contains examples of manifests. They can be used to try out the controller.

```
# Create 
kubectl apply -f config/samples/subnetmachinerequest_v1alpha1_subnetmachinerequest.yaml
TODO
```

### Cleanup

`make undeploy`

### Testing

```
# Go to controller directory
cd controllers

# Run tests
go test . -v -ginkgo.v

```

## Project created with operator SDK (go - kubebuilder)

Steps to reproduce: 
- init ` operator-sdk init --domain onmetal.de --repo github.com/onmetal/k8s-subnet-machine-request`
- core `operator-sdk create api --group subnetmachinerequest --version v1alpha1 --kind SubnetMachineRequest --resource --controller`

## Diagram

![Diagram](./docs/subnetmr.jpg)
