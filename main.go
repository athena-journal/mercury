package main

import (
	discovery "github.com/athena-journal/mercury/discovery"
	utils "github.com/athena-journal/mercury/utils"
	"github.com/gin-gonic/gin"
)

// Constants
const PUBLIC_NAMESPACE = "athena-public"
const PRIVATE_NAMESPACE = "athena-private"

func main() {
	go discovery.GetAllServices(PUBLIC_NAMESPACE, PRIVATE_NAMESPACE)
	router := gin.Default()

	// proxyPath is the path to the service, services get routed through here for service discovery
	// Only support these HTTP Verbs
	router.GET("/*proxyPath", utils.EndpointHandler)
	router.POST("/*proxyPath", utils.EndpointHandler)
	router.PUT("/*proxyPath", utils.EndpointHandler)
	router.DELETE("/*proxyPath", utils.EndpointHandler)

	// Serve Gin on Port 8080
	router.Run(":8080")
}
