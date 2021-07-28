package server

import (
	"net/http"

	"github.com/Ibra1994/go-start/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("v1")
	{
		organizationController := new(controllers.OrganizationController)
		v1.GET("organization", organizationController.Get)
		v1.GET("organization/:id", organizationController.Show)
		v1.POST("organization", organizationController.Create)
	}

	return r
}
