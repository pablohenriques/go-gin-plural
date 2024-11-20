package main

import (
	"errors"
	"gin-course-plural/employee"
	"github.com/gin-contrib/gzip"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	//r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(myErrorLogger)
	r.Use(gin.CustomRecovery(myRecoveryFunc))

	registerRoutes(r)

	r.Run()

}

func registerRoutes(r *gin.Engine) {

	tryToGetEmployee := func(c *gin.Context, employeeIDRaw string) (*employee.Employee, bool) {
		employeeID, err := strconv.Atoi(employeeIDRaw)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return nil, false
		}
		emp, err := employee.Get(employeeID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return nil, false
		}
		return emp, true
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/employees")
	})

	r.GET("/employees", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", employee.GetAll())
	})

	r.GET("/employees/:employeeID", func(c *gin.Context) {
		employeeIDRaw := c.Param("employeeID")

		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			c.HTML(http.StatusOK, "employee.tmpl", *emp)
		}
	})

	r.POST("/employees/:employeeID", func(c *gin.Context) {
		var timeoff employee.TimeOff
		err := c.ShouldBind(&timeoff)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested

		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			c.Redirect(http.StatusFound, "/employees/"+employeeIDRaw)
		}
	})

	r.GET("/errors", func(c *gin.Context) {
		err := &gin.Error{
			Err:  errors.New("something went horribly wrong"),
			Type: gin.ErrorTypeRender | gin.ErrorTypePublic,
			Meta: "this error was intentional",
		}
		c.Error(err)
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("a Go program should almostnever  call 'panic'")
	})

	g := r.Group("/api/employees", Benchmark)
	{
		g.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, employee.GetAll())
		})
		g.GET("/:employeeID", func(c *gin.Context) {
			employeeIDRaw := c.Param("employeeID")
			if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
				c.JSON(http.StatusOK, *emp)
			}
		})
		g.POST("/:employeeID", func(c *gin.Context) {
			var timeoff employee.TimeOff
			err := c.ShouldBind(&timeoff)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			timeoff.Type = employee.TimeoffTypePTO
			timeoff.Status = employee.TimeoffStatusRequested

			employeeIDRaw := c.Param("employeeID")
			if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
				emp.TimeOff = append(emp.TimeOff, timeoff)
				c.JSON(http.StatusOK, *emp)
			}
		})
	}

	r.Static("/public", "./public")
}

var Benchmark gin.HandlerFunc = func(c *gin.Context) {
	t := time.Now()
	c.Next()
	elapsed := time.Since(t)
	log.Print("Time to process", elapsed)
}

var myErrorLogger gin.HandlerFunc = func(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		log.Print(map[string]any{
			"Err":  err.Error(),
			"Type": err.Type,
			"Meta": err.Meta,
		})
	}
}

var myRecoveryFunc gin.RecoveryFunc = func(c *gin.Context, err any) {
	log.Print("Custom recovery functions can be used to add fine=grained control over recovery strategies", err)
}
