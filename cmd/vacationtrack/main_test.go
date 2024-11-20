package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiEmployees(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/employees/", nil)

	r := gin.Default()
	registerRoutes(r)

	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
	t.Log(rec.Body.String())
}
