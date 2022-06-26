package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func sendRequest(method string, url string, c *gin.Context) {
	var req *http.Request
	var err error

	// If it's a GET || Delete Request we don't have a body, otherwise we have a body
	if method != http.MethodGet && method != http.MethodDelete {
		rawBody, _ := ioutil.ReadAll(c.Request.Body)
		req, err = http.NewRequest(method, url, strings.NewReader(string(rawBody)))
		if err != nil {
			c.JSON(500, gin.H{
				"message": "500 Internal Server Error",
			})
			return
		}
	} else {
		req, err = http.NewRequest(method, url, http.NoBody)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "500 Internal Server Error",
			})
			return
		}
	}

	client := &http.Client{}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "500 Internal Server Error",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8") // Content Type Header
	req.Header.Add("Connection", "close")                             // Close Connection after request is done

	// Send Request
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "500 Internal Server Error",
		})
		return
	}

	// Stream Response to the Gin Context
	handleResponse(resp, c)
}
