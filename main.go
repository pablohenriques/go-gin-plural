package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type TimeOffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"-" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"-"`
}

func main() {
	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		var timeoffRequest TimeOffRequest
		if err := c.ShouldBind(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	apiGroup := router.Group("/api")
	apiGroup.POST("/timeoff", func(c *gin.Context) {
		var timeoffRequest TimeOffRequest
		if err := c.ShouldBind(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	log.Fatal(router.Run(":3000"))
}
