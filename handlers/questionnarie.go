package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary				Get Questionnaires
//	@Description.markdown	questionnaires.get
//	@Tags					Questionnaire
//	@Produce				json
//	@Success				200	{object}	[]Questionnaire
//	@Router					/questionnaires [get]
func (h *QuestionnaireHandler) GetQuestionnaires(c *gin.Context) {
	questionnaires, err := h.DB.Questionnaire.
		Query().
		WithQuestions().
		WithQuestionnaireResponses(func(qr *ent.QuestionnaireResponseQuery) {
			qr.WithUser().WithAnswers().All(c.Request.Context())
		}).
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionnaires)
}

//	@Summary				Get Questionnaire
//	@Description.markdown	questionnaire.get
//	@Tags					Questionnaire
//	@Produce				json
//	@Param					id	path		string	true	"The questionnaire's ID"
//	@Success				200	{object}	Questionnaire
//	@Router					/questionnaires/{id} [get]
func (h *QuestionnaireHandler) GetQuestionnaire(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	questionnaireNode, err := h.DB.Questionnaire.
		Query().
		Where(questionnaire.ID(id)).
		WithQuestions().
		WithQuestionnaireResponses(func(qr *ent.QuestionnaireResponseQuery) {
			qr.WithUser().WithAnswers().All(c.Request.Context())
		}).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionnaireNode)
}

//	@Summary				Create Questionnaire
//	@Description.markdown	questionnaire.post
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					questionnaire	body		SingleQuestionnaireResponse	true	"The questionnaire to be created."
//	@Success				200				{object}	Questionnaire
//	@Router					/questionnaires [post]
func (h *QuestionnaireHandler) CreateQuestionnaire(c *gin.Context) {
	var questionnaireBody QuestionnaireBody
	if err := c.ShouldBindJSON(&questionnaireBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := json.MarshalIndent(questionnaireBody, "", "  ")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Print(string(out))

	questionnaireNode, err := h.DB.Questionnaire.
		Create().
		SetName(questionnaireBody.Name).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, question := range questionnaireBody.Questions {
		_, err := h.DB.Question.
			Create().
			SetBody(question.Body).
			SetType(question.Type).
			SetQuestionnaire(questionnaireNode).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, questionnaireNode)
}

//	@Summary				Delete Questionnaire
//	@Description.markdown	questionnaire.delete
//	@Tags					Questionnaire
//	@Produce				json
//	@Param					id	path	string	true	"The questionnaire's ID."
//	@Success				200
//	@Router					/questionnaires/{id} [delete]
func (h *QuestionnaireHandler) DeleteQuestionnaire(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Questionnaire.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary				Create Question
//	@Description.markdown	question.post
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					id			path		string			true	"The questionnaire's ID."
//	@Param					question	body		QuestionBody	true	"The question to be created."
//	@Success				200			{object}	ent.Question
//	@Failure				400			{object}	ent.Question
//	@Router					/questionnaires/{id}/new/question [post]
func (h *QuestionnaireHandler) CreateQuestion(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var questionBody ent.Question
	if err := c.ShouldBindJSON(&questionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questionNode, err := h.DB.Question.
		Create().
		SetBody(questionBody.Body).
		SetType(questionBody.Type).
		SetQuestionnaireID(id).
		Save(c.Request.Context())

	if err != nil {
		errH := gin.H{"error": err.Error()}

		if ent.IsConstraintError(err) {
			c.JSON(http.StatusBadRequest, errH)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionNode)
}

//	@Summary				Create Response
//	@Description.markdown	response.post
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					id			path		string						true	"The questionnaire's ID."
//	@Param					response	body		QuestionnaireResponseBody	true	"The response to be created."
//	@Success				200			{object}	QuestionnaireResponse
//	@Router					/questionnaires/{id}/new/response [post]
func (h *QuestionnaireHandler) CreateResponse(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseBody QuestionnaireResponseBody
	if err := c.ShouldBindJSON(&responseBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseNode, err := h.DB.QuestionnaireResponse.
		Create().
		SetUserID(responseBody.UserId).
		SetQuestionnaireID(id).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, answer := range responseBody.Answers {
		_, err := h.DB.Answer.
			Create().
			SetBody(answer.Body).
			SetQuestionID(answer.QuestionId).
			SetQuestionnaireResponse(responseNode).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responseNode)
}
