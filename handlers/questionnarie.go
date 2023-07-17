package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/gin-gonic/gin"
)

type QuestionnaireHandler struct {
	DB *ent.Client
}

type Question struct {
	Body string `json:"body"`
	Type string `json:"type"`
}

type QuestionnaireBody struct {
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type Answer struct {
	QuestionId int    `json:"question_id"`
	Body       string `json:"body"`
}

type ResponseBody struct {
	UserId  string   `json:"user_id"`
	Answers []Answer `json:"answers"`
}

func (h *QuestionnaireHandler) GetAllQuestionnaires(c *gin.Context) {
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

func (h *QuestionnaireHandler) GetQuestionnaireById(c *gin.Context) {
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

	questions := make([]*ent.Question, 0, len(questionnaireBody.Questions))
	for _, question := range questionnaireBody.Questions {
		questionNode, err := h.DB.Question.
			Create().
			SetBody(question.Body).
			SetType(question.Type).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		questions = append(questions, questionNode)
	}

	questionnaire, err := h.DB.Questionnaire.
		Create().
		SetName(questionnaireBody.Name).
		AddQuestions(questions...).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionnaire)
}

func (h *QuestionnaireHandler) CreateResponse(c *gin.Context) {
	questionnaireId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseBody ResponseBody
	if err := c.ShouldBindJSON(&responseBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Print(string(out))

	answerNodes := make([]*ent.Answer, 0, len(responseBody.Answers))
	for _, answer := range responseBody.Answers {
		answerNode, err := h.DB.Answer.
			Create().
			SetBody(answer.Body).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		answerNodes = append(answerNodes, answerNode)
	}

	response, err := h.DB.QuestionnaireResponse.
		Create().
		AddAnswers(answerNodes...).
		AddQuestionnaireIDs(questionnaireId).
		AddUserIDs(responseBody.UserId).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *QuestionnaireHandler) DeleteQuestionnaireById(c *gin.Context) {
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
