package handlers

import (
	"net/http"
	"strconv"

	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/gin-gonic/gin"
)

// GET		/questions
func (h *QuestionHandler) GetQuestions(c *gin.Context) {
	questions, err := h.DB.Question.
		Query().
		WithQuestionnaire().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// GET		/questions/:id
func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	questions, err := h.DB.Question.
		Query().
		Where(question.ID((id))).
		WithQuestionnaire().
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// DELETE	/questions/:id
func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Question.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
// 	questionnaireId, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var questionBody QuestionBody
// 	if err := c.ShouldBindJSON(&questionBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	out, err := json.MarshalIndent(questionBody, "", "  ")
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	log.Print(string(out))

// 	answerNodes := make([]*ent.Answer, 0, len(questionBody.Answers))
// 	for _, answer := range questionBody.Answers {
// 		answerNode, err := h.DB.Answer.
// 			Create().
// 			SetBody(answer.Body).
// 			Save(c.Request.Context())

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 			return
// 		}

// 		answerNodes = append(answerNodes, answerNode)
// 	}

// 	question, err := h.DB.QuestionnaireQuestion.
// 		Create().
// 		AddAnswers(answerNodes...).
// 		AddQuestionnaireIDs(questionnaireId).
// 		AddUserIDs(questionBody.UserId).
// 		Save(c.Request.Context())

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, question)
// }
