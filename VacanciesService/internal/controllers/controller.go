package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) get_all(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, map[string]string{
		"message:": "hello world",
	})
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/vacancies/v1/all", c.get_all)

	return router
}
