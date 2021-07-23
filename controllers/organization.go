package controllers

import (
	"net/http"

	"github.com/Ibra1994/go-start/models"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct{}

var model models.OrganizationModel

func (ctrl OrganizationController) Get(c *gin.Context) {
	organizations := model.Get()
	c.JSON(http.StatusOK, organizations)
}

func (ctrl OrganizationController) Show(c *gin.Context) {

	id := c.Param("id")
	organization := model.Show(id)
	c.JSON(http.StatusOK, organization)
}

func (ctrl OrganizationController) Create(c *gin.Context) {
	var create models.Organization
	c.ShouldBindJSON(&create)

	organization := model.Create(create)

	c.JSON(http.StatusOK, organization)
}

func (ctrl OrganizationController) Update(c *gin.Context) {

}

func (ctrl OrganizationController) Delete(c *gin.Context) {

}
