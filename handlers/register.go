package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

// @Summary				Register an User
// @Description.markdown	register.post
// @Tags					Registration
// @Accept					json
// @Produce				json
// @Param					user	body		types.RegisterData	true	"The registration data."
// @Success				200		{object}	types.RegisterResponse
// @Router					/register [post]
func (h *Handler) Register(c *gin.Context) {

	// STEP 0:	Check if a MyCard is available
	myCardNode, err := h.DB.MyCard.Query().
		Where(mycard.Not(mycard.HasRecipient())).
		First(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// STEP 1:	Create a new user
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
		SetMedicalHistory(user.MedicalHistory(body.User.MedicalHistory)).
		SetMedicationStatus(user.MedicationStatus(body.User.MedicationStatus)).
		SetDementedAmongDirectRelatives(body.User.DementedAmongDirectRelatives).
		SetHeadInjuryExperience(body.User.HeadInjuryExperience).
		SetEarCondition(user.EarCondition(body.User.EarCondition)).
		SetEyesightCondition(user.EyesightCondition(body.User.EyesightCondition)).
		SetSmokingHabit(user.SmokingHabit(body.User.SmokingHabit)).
		AddMycards(myCardNode). // Assign a MyCard
		Save(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// STEP 2:	Respond to registration questionnaire
	raw_questionnaire_id := body.Response.QuestionnaireId
	if raw_questionnaire_id != "88888888-8888-4888-8888-888888888888" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "you should submit the registration questionnaire instead"},
		)
	}

	responseNode, err := h.RespondQuestionnaire(
		c.Request.Context(),
		userNode.ID,
		raw_questionnaire_id,
		body.Response.Answers,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":     userNode,
		"mycard":   myCardNode,
		"response": responseNode,
	})

	// TODO
	// STEP	3:	Send notification
}
