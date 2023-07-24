package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary				Register an User
//	@Description.markdown	register.post
//	@Tags					Registration
//	@Accept					json
//	@Produce				json
//	@Param					user	body		types.RegisterData	true	"The registration data."
//	@Success				200		{object}	types.RegisterResponse
//	@Router					/register [post]
func (h *RegisterHandler) Register(c *gin.Context) {
	var body types.RegisterData
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Print(string(out))

	userNode, err := h.DB.User.
		Create().
		SetID(body.User.ID).
		SetBirthYear(body.User.BirthYear).
		SetHeight(body.User.Height).
		SetWeight(body.User.Weight).
		SetGender(user.Gender(body.User.Gender)).
		SetEducationLevel(user.EducationLevel(body.User.EducationLevel)).
		SetOccupation(user.Occupation(body.User.Occupation)).
		SetMarriage(user.Marriage(body.User.Marriage)).
		SetMedicalHistory(body.User.MedicalHistory).
		SetMedicationStatus(body.User.MedicationStatus).
		SetDementedAmongDirectRelatives(body.User.DementedAmongDirectRelatives).
		SetHeadInjuryExperience(body.User.HeadInjuryExperience).
		SetEarCondition(user.EarCondition(body.User.EarCondition)).
		SetEyesightCondition(user.EyesightCondition(body.User.EyesightCondition)).
		SetSmokingHabit(user.SmokingHabit(body.User.SmokingHabit)).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	questionnaire_id, err := uuid.Parse(c.Param(body.Response.QuestionnaireId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseNode, err := h.DB.QuestionnaireResponse.
		Create().
		SetUser(userNode).
		SetQuestionnaireID(questionnaire_id).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	for _, answer := range body.Response.Answers {
		_, err := h.DB.Answer.
			Create().
			SetBody(answer.Body).
			SetQuestionID(answer.QuestionId).
			SetQuestionnaireResponse(responseNode).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user":     userNode,
		"response": responseNode,
	})

	// TODO: Send notification
}
