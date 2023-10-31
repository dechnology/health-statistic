package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Description	Datatype of health status
type HealthStatus struct {
	Message string `json:"message" example:"Hello, this is an example message!"` // Health message
}

var (
	// failureMessage = "Hello from Dechnology! The database is disconnected!"
	successMessage = "Hello from Dechnology! The server is functioning!"
)

//	@Summary		Health Check
//	@Description	A health checking endpoint to make sure the server is not dead.
//	@Tags			Health
//
//	@Security		BearerAuth
//
//	@Success		200	{object}	HealthStatus
//	@Router			/health_check [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	// if err := h.DB.SQLDB.Ping(); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": failureMessage,
	// 	})
	// 	return
	// }
	//
	c.JSON(http.StatusOK, gin.H{
		"message": successMessage,
	})
}
