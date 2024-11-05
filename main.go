package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "./index.html")

	router.GET("/tale_of_two_cities", func(c *gin.Context) {
		f, err := os.Open("./a_tale_of_two_cities.txt")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Data(http.StatusOK, "text/plain", data)

	})

	router.GET("/great_expectations", func(c *gin.Context) {
		f, err := os.Open("./great_expectations.txt")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.DataFromReader(http.StatusOK,
			fi.Size(),
			"text/plain",
			f,
			map[string]string{
				"Content-Disposition": "attachment;filename=great_expetations.txt",
			},
		)
	})

	log.Fatal(router.Run(":3000"))
}
