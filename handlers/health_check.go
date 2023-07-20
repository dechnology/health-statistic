package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckStatus struct {
	Message string `json:"message"`
}

var (
	successMessage = "Hello from Dechnology! The server is functioning!"
)

//	@Summary		Health Check
//	@Description	A health checking api to make sure the server is not dead.
//	@Tags			Health
//	@Success		200	object	HealthCheckStatus
//	@Router			/health_check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckStatus{Message: successMessage})
}
