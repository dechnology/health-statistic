package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Health Check
// @Description A health checking api to make sure the server is not dead.
// @Tags        Response
// @Success 	200
// @Router      /health_check [get]
func HealthCheck(c *gin.Context) {
	responses, err := h.DB.QuestionnaireResponse.
		Query().
		WithAnswers().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}
