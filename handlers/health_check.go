package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     Health Check
// @Description A health checking api to make sure the server is not dead.
// @Tags        Response
// @Success 	200
// @Router      /health_check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
