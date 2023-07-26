package handlers

import (
	"context"
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) RespondQuestionnaire(
	ctx context.Context,
	user_id, raw_questionnaire_id string,
	answers []*types.BaseAnswer,
) (*ent.QuestionnaireResponse, error) {
	questionnaire_id, err := uuid.Parse(raw_questionnaire_id)
	if err != nil {
		return nil, err
	}

	responseNode, err := h.DB.QuestionnaireResponse.
		Create().
		SetUserID(user_id).
		SetQuestionnaireID(questionnaire_id).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	for _, answer := range answers {
		_, err := h.DB.Answer.
			Create().
			SetBody(*answer.Body).
			SetQuestionID(answer.QuestionId).
			SetQuestionnaireResponse(responseNode).
			Save(ctx)

		if err != nil {
			return nil, err
		}
	}

	return responseNode, nil
}

//	@Summary				Get Responses
//	@Description.markdown	responses.get
//	@Tags					Response
//	@Produce				json
//	@Success				200	{object}	[]types.ResponseWithQuestionnaire
//	@Router					/responses [get]
func (h *Handler) GetResponses(c *gin.Context) {
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
//	@Success				200	{object}	types.ResponseWithQuestionnaire
//	@Router					/responses/{id} [get]
func (h *Handler) GetResponse(c *gin.Context) {
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
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
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
func (h *Handler) DeleteResponse(c *gin.Context) {
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
