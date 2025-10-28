package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thebiatriz/golang/statuscheck/internal/check"
)

func main() {
	router := gin.Default()

	router.GET("/check", check.CheckMultipleURLs)

	router.Run(":8080")
}