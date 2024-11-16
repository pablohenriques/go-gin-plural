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
		c.Stream(streamer(f))

	})

	log.Fatal(router.Run(":3000"))
}

func streamer(r io.Reader) func(io.Writer) bool {
	return func(step io.Writer) bool {
		for {
			buf := make([]byte, 4*2^10)
			if _, err := r.Read(buf); err == nil {
				_, err := step.Write(buf)
				return err == nil
			} else {
				return false
			}
		}
	}
}
