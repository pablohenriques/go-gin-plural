package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/query/*rest", func(c *gin.Context) {
		username := c.Query("username")
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := c.QueryArray("month")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})

	})

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		date := c.PostForm("date")
		amount := c.PostForm("amount")
		username := c.DefaultPostForm("username", "adent")

		c.JSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		})

	})

	log.Fatal(router.Run(":3000"))
}
