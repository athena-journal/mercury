package discovery

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Services
type Service struct {
	Name string
	Port int32
}

var PublicServices = map[string]Service{}
var PrivateServices = map[string]Service{}

/*
  Ran as a Go Routine, Gets all Services from the cluster and stores them in a map
*/
func GetAllServices(PUBLIC_NAMESPACE string, PRIVATE_NAMESPACE string) {
	ticker := time.NewTicker(10 * time.Second) // Check for services every 10s
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				getServices(PUBLIC_NAMESPACE, PRIVATE_NAMESPACE)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

/*
  Gets all Kubernetes services within the cluster
*/
func getServices(PUBLIC_NAMESPACE, PRIVATE_NAMESPACE string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		_ = fmt.Errorf("error getting cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		_ = fmt.Errorf("error creating clientset: %v", err)
	}

	publicServiceList, err := clientset.CoreV1().Services(PUBLIC_NAMESPACE).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		_ = fmt.Errorf("error getting services: %v", err)
	}

	privateServiceList, err := clientset.CoreV1().Services(PRIVATE_NAMESPACE).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		_ = fmt.Errorf("error getting services: %v", err)
	}

	for _, service := range publicServiceList.Items {
		PublicServices[service.Name] = Service{
			Name: service.Name,
			Port: service.Spec.Ports[0].Port,
		}
	}

	for _, service := range privateServiceList.Items {
		PublicServices[service.Name] = Service{
			Name: service.Name,
			Port: service.Spec.Ports[0].Port,
		}
	}
}
