package handlers

import (
	"net/http"
	"strconv"

	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/gin-gonic/gin"
)

// GET		/responses
func (h *ResponseHandler) GetResponses(c *gin.Context) {
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

// GET		/responses/:id
func (h *ResponseHandler) GetResponse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses, err := h.DB.QuestionnaireResponse.
		Query().
		Where(questionnaireresponse.ID((id))).
		WithAnswers().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// DELETE	/responses/:id
func (h *ResponseHandler) DeleteResponse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.QuestionnaireResponse.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
