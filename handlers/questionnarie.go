package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetQuestionnaireById(
	ctx context.Context,
	userId string,
	questionnaireId uuid.UUID,
) (*ent.Questionnaire, error) {
	questionnaireNode, err := h.DB.Questionnaire.
		Query().
		Where(questionnaire.ID(questionnaireId)).
		WithQuestions(func(qq *ent.QuestionQuery) {
			qq.WithChoices().All(ctx)
		}).
		WithQuestionnaireResponses(func(qrq *ent.QuestionnaireResponseQuery) {
			qrq.Where(questionnaireresponse.HasUserWith(user.ID(userId))).
				WithAnswers(func(aq *ent.AnswerQuery) {
					aq.WithChosen()
				})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return questionnaireNode, nil
}

func (h *Handler) AppendQuestions(
	ctx context.Context,
	questionnaireId uuid.UUID,
	questions []types.BaseQuestion,
) error {
	n_questions, err := h.DB.Questionnaire.Query().QueryQuestions().Count(ctx)
	if err != nil {
		return err
	}

	for i, questionData := range questions {
		questionNode, err := h.DB.Question.Create().
			SetType(question.Type(questionData.Type)).
			SetBody(questionData.Body).
			SetOrder(n_questions + i).
			SetQuestionnaireID(questionnaireId).
			Save(ctx)
		if err != nil {
			return err
		}

		for i, choice := range questionData.Choices {
			_, err := h.DB.Choice.Create().
				SetBody(choice).
				SetQuesion(questionNode).
				SetOrder(i).
				Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//	@Summary				Get Questionnaires
//	@Description.markdown	questionnaires.get
//	@Tags					Questionnaire
//	@Produce				json
//	@Success				200	{object}	[]types.QuestionnaireDetails
//	@Router					/questionnaires [get]
func (h *Handler) GetQuestionnaires(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	questionnaireNodes, err := h.DB.Questionnaire.
		Query().
		WithQuestions(func(qq *ent.QuestionQuery) {
			qq.WithChoices().All(c.Request.Context())
		}).
		WithQuestionnaireResponses(func(qrq *ent.QuestionnaireResponseQuery) {
			qrq.Where(questionnaireresponse.HasUserWith(user.ID(*userId))).
				WithAnswers(func(aq *ent.AnswerQuery) {
					aq.WithChosen().All(c.Request.Context())
				}).
				All(c.Request.Context())
		}).
		All(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, questionnaireNodes)
}

//	@Summary				Get Registration Questionnaire
//	@Description.markdown	registration_questionnaire.get
//	@Tags					Questionnaire
//	@Produce				json
//	@Success				200	{object}	types.QuestionnaireDetails
//	@Router					/questionnaires/registration [get]
func (h *Handler) GetRegistrationQuestionnaire(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	questionnaireNode, err := h.GetQuestionnaireById(
		c.Request.Context(),
		*userId,
		h.RegistrationQuestionnaireId,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, questionnaireNode)
}

//	@Summary				Get Questionnaire
//	@Description.markdown	questionnaire.get
//	@Tags					Questionnaire
//	@Produce				json
//	@Param					id	path		string	true	"The questionnaire's ID"
//	@Success				200	{object}	types.QuestionnaireDetails
//	@Router					/questionnaires/{id} [get]
func (h *Handler) GetQuestionnaire(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	questionnaireId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	questionnaireNode, err := h.GetQuestionnaireById(
		c.Request.Context(),
		*userId,
		questionnaireId,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, questionnaireNode)
}

//	@Summary				Create Questionnaire
//	@Description.markdown	questionnaire.post
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					questionnaire	body		types.BaseQuestionnaire	true	"The questionnaire to be created."
//	@Success				200				{object}	ent.Questionnaire
//	@Router					/questionnaires [post]
func (h *Handler) CreateQuestionnaire(c *gin.Context) {
	var body types.BaseQuestionnaire
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(out))

	questionnaireNode, err := h.DB.Questionnaire.
		Create().
		SetName(body.Name).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.AppendQuestions(
		c.Request.Context(), questionnaireNode.ID, body.Questions,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
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
func (h *Handler) DeleteQuestionnaire(c *gin.Context) {
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

//	@Summary				Create Questionnaire Response
//	@Description.markdown	questionnaire_responses.post
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					id			path		string				true	"The questionnaire's ID."
//	@Param					response	body		types.BaseResponse	true	"The response to be created."
//	@Success				200			{object}	ent.QuestionnaireResponse
//	@Router					/questionnaires/{id}/responses [post]
func (h *Handler) CreateQuestionnaireResponse(c *gin.Context) {
	questionnaireId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	var body types.ResponseWithUserId
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseNode, err := h.RespondQuestionnaire(
		c.Request.Context(),
		body.UserId,
		questionnaireId,
		body.Answers,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, responseNode)
}

//	@Summary				Get Questionnaire Responses
//	@Description.markdown	questionnaire_responses.get
//	@Tags					Questionnaire
//	@Accept					json
//	@Produce				json
//	@Param					id	path		string	true	"The questionnaire's ID."
//	@Success				200	{object}	[]ent.QuestionnaireResponse
//	@Router					/questionnaires/{id}/responses [get]
func (h *Handler) GetQuestionnaireResponses(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	responseNodes, err := h.DB.QuestionnaireResponse.Query().
		Where(questionnaireresponse.HasQuestionnaireWith(questionnaire.ID(id))).
		WithAnswers(func(q *ent.AnswerQuery) {
			q.WithChosen().All(c.Request.Context())
		}).
		All(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, responseNodes)
}
