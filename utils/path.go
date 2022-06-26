package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

/*
  Extracts the service and path from the request
*/
func extractPath(c *gin.Context) (service string, path string) {
	proxyPath := c.Param("proxyPath")
	service = strings.Split(proxyPath, "/")[1]
	path = proxyPath[len(service)+1:] // Get the URL path of the service

	return service, path
}
