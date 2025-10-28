package check

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckRoute(c *gin.Context) {
	url := c.Query("url")
	if strings.TrimSpace(url) == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Url está vazia"})
		return
	}

	status, err := checkSite(url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := gin.H{
		"status": status,
		"url":    url,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func checkSite(urlToCheck string) (string, error) {
	resp, err := http.Head(urlToCheck)

	if err != nil {
		return "ERROR", err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "UP", nil
	} else {
		return "DOWN", err
	}
}