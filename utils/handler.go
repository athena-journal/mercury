package handler

import (
	"net/http"
	"strconv"

	"github.com/athena-journal/mercury/discovery"
	"github.com/gin-gonic/gin"
)

/*
  Handles the request to the service
*/
func EndpointHandler(c *gin.Context) {
	// Get the service and url from the request
	service, path := extractPath(c)
	if service == "" || path == "" {
		c.JSON(404, gin.H{
			"message": "404 Not Found",
		})
		return
	}

	if _, ok := discovery.PublicServices[service]; ok {
		// Just call the service, handles load balancing and other things for us
		// All services we want to be publicly accessible are stored in the 'athena-public' namespace
		url := "http://" + discovery.PublicServices[service].Name + ".athena-public" + ":" + strconv.Itoa(int(discovery.PublicServices[service].Port)) + path

		switch {
		case c.Request.Method == "GET":
			sendRequest(http.MethodGet, url, c)
			return
		case c.Request.Method == "POST":
			sendRequest(http.MethodPost, url, c)
			return
		case c.Request.Method == "PUT":
			sendRequest(http.MethodPut, url, c)
			return
		case c.Request.Method == "DELETE":
			sendRequest(http.MethodDelete, url, c)
		default:
			// Don't handle any other HTTP verbs
			c.JSON(404, gin.H{
				"message": "404 Not Found",
			})
			return
		}
	}

	// If we get here, the service is not public
	c.JSON(404, gin.H{
		"message": "404 Not Found",
	})
}
