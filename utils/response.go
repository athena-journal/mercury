package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
  Handles Response back from Microservice(s)
*/
func handleResponse(resp *http.Response, c *gin.Context) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "500 Internal Server Error",
		})
	}

	c.Data(resp.StatusCode, "application/json", body)
}
