package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/gin-gonic/gin"
)

// GET		/questionnaires
func (h *QuestionnaireHandler) GetQuestionnaires(c *gin.Context) {
	questionnaires, err := h.DB.Questionnaire.
		Query().
		WithQuestions().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionnaires)
}

// GET		/questionnaires/:id
func (h *QuestionnaireHandler) GetQuestionnaire(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	questionnaireNode, err := h.DB.Questionnaire.
		Query().
		Where(questionnaire.ID(id)).
		WithQuestions().
		WithQuestionnaireResponses(func(qr *ent.QuestionnaireResponseQuery) {
			qr.WithAnswers()
		}).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionnaireNode)
}

// POST		/questionnaires
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

// DELETE	/questionnaires/:id
func (h *QuestionnaireHandler) DeleteQuestionnaire(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

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

// GET		/questionnaires/:id/questions
func (h *QuestionnaireHandler) GetQuestions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	questions, err := h.DB.Question.
		Query().
		Where(question.HasQuestionnaireWith(questionnaire.ID(id))).
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// POST		/questionnaires/:id/questions
func (h *QuestionnaireHandler) CreateQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var questionBody Question
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionNode)
}

// GET		/questionnaires/:id/responses
func (h *QuestionnaireHandler) GetResponses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses, err := h.DB.QuestionnaireResponse.
		Query().
		Where(questionnaireresponse.HasQuestionnaireWith(questionnaire.ID(id))).
		WithAnswers().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// POST		/questionnaires/:id/responses
func (h *QuestionnaireHandler) CreateResponse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseBody ResponseBody
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
