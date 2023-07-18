package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Get Questions
// @Description Get all questions from the database.
// @Tags        Question
// @Produce     json
// @Success 	200 {object} []Question
// @Router      /questions [get]
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

// @Summary     Get Question
// @Description Get a question by ID.
// @Tags        Question
// @Produce     json
// @Param		id path string true "The question's ID"
// @Success 	200 {object} Question
// @Router      /questions/{id} [get]
func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

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

// @Summary     Delete Question
// @Description Delete a question by ID
// @Tags        Question
// @Produce     json
// @Param		id path string true "The question's ID."
// @Success 	200
// @Router      /questions/{id} [delete]
func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

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
