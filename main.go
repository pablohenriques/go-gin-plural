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
	adminGroup := router.Group("/admin")

	adminGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to administer roles")
	})

	adminGroup.GET("/roles", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to administer policies")
	})

	adminGroup.GET("/policies", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to administer users")
	})

	log.Fatal(router.Run(":3000"))
}
