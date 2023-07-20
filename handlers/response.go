package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary				Get Responses
//	@Description.markdown	responses.get
//	@Tags					Response
//	@Produce				json
//	@Success				200	{object}	[]QuestionnaireResponse
//	@Router					/responses [get]
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

//	@Summary				Get Response
//	@Description.markdown	response.get
//	@Tags					Response
//	@Produce				json
//	@Param					id	path		string	true	"The response's ID"
//	@Success				200	{object}	SingleQuestionnaireResponse
//	@Router					/responses/{id} [get]
func (h *ResponseHandler) GetResponse(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses, err := h.DB.QuestionnaireResponse.
		Query().
		Where(questionnaireresponse.ID((id))).
		WithQuestionnaire().
		WithAnswers().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

//	@Summary				Delete Response
//	@Description.markdown	response.delete
//	@Tags					Response
//	@Produce				json
//	@Param					id	path	string	true	"The response's ID."
//	@Success				200
//	@Router					/responses/{id} [delete]
func (h *ResponseHandler) DeleteResponse(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

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
