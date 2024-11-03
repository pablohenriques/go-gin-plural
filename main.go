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
	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})
	router.POST("/employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New request POSTed successfully!")
	})
	log.Fatal(router.Run(":3000"))
}
