package main

import (
	"gin-course-plural/employee"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
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

	r.Static("/public", "./public")
}
