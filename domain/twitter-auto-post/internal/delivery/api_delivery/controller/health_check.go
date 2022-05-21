package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (HC *HealthCheckController) Alive(c *gin.Context) {
	resp := gin.H{
		"status":  http.StatusText(http.StatusOK),
		"service": "name-service",
	}

	c.JSON(http.StatusOK, resp)
}
