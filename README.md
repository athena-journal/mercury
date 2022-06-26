# Mercury - Edge Proxy Service Built for Kubernetes

Mercury is a proxy service used to proxy requests to an internal service over HTTP within your Kubernetes cluster.

### Context

Mercury runs on the [Gin](https://github.com/gin-gonic/gin) Go Framework to facilitate high throughput requests and is built on top of the [Kubernetes](https://kubernetes.io/) API.

It provides high-level abstractions to improve your workflow and reduce the amount of Kubernetes YAML files you need to maintain. By automatically routing requests to the correct service, you can focus on the core functionality of your application.

### Configuration

Mercury can easily be configured by modifying the `main.go` file in the root of this project.

```go
// Constants
const PUBLIC_NAMESPACE = "athena-public" // Namespace of Services you want to be able to access from the edge
const PRIVATE_NAMESPACE = "athena-private" // Namespace of internal services used for authentication, Mercury will not allow access to services in this namespace from the edge without prior configuration
const POLL_K8_API_SECONDS = 10 // How often to poll K8s API for changes to services
```

### First Time installation

*<b>Make sure you have Go & Docker installed locally before starting this process</b>*

Switch to the Kubernetes cluster you want to deploy Mercury to

```bash
kubectl config use-context <cluster-name>
``` 

Then create the `athena-edge` namespace

```bash
kubectl create namespace athena-edge
```

Finally just run `make ship` to build and deploy Mercury on your cluster!
