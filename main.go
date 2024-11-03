package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()
	router.GET("/employees/:username/*rest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.Param("username"),
			"rest":     c.Param("rest"),
		})
	})
	log.Fatal(router.Run(":3000"))
}
