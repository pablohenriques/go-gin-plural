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
	router.StaticFile("/", "./public/index.html")
	router.Static("/public", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(f)))
	log.Fatal(router.Run(":3000"))
}
