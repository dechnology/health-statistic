package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Description	Datatype of health status
type HealthStatus struct {
	Message string `json:"message"` // Health message
}

var (
	successMessage = "Hello from Dechnology! The server is functioning!"
)

//	@Summary		Health Check
//	@Description	A health checking endpoint to make sure the server is not dead.
//	@Tags			Health
//	@Success		200	{object}	HealthStatus
//	@Router			/health_check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthStatus{Message: successMessage})
}
