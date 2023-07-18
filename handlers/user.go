package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *ent.Client
}

type UserBody struct {
	ID                           string  `json:"id"`
	BirthYear                    int     `json:"birth_year"`
	Height                       float64 `json:"height"`
	Weight                       float64 `json:"weight"`
	Gender                       string  `json:"gender"`
	EducationLevel               string  `json:"education_level"`
	Occupation                   string  `json:"occupation"`
	Marriage                     string  `json:"marriage"`
	MedicalHistory               string  `json:"medical_history"`
	MedicationStatus             string  `json:"medication_status"`
	DementedAmongDirectRelatives bool    `json:"demented_among_direct_relatives"`
	HeadInjuryExperience         bool    `json:"head_injury_experience"`
	EarCondition                 string  `json:"ear_condition"`
	EyesightCondition            string  `json:"eyesight_condition"`
	SmokingHabit                 string  `json:"smoking_habit"`
}

// GET		/users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.DB.User.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GET		/users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	user, err := h.DB.User.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// POST		/users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var body UserBody
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
		SetID(body.ID).
		SetBirthYear(body.BirthYear).
		SetHeight(body.Height).
		SetWeight(body.Weight).
		SetGender(user.Gender(body.Gender)).
		SetEducationLevel(user.EducationLevel(body.EducationLevel)).
		SetOccupation(user.Occupation(body.Occupation)).
		SetMarriage(user.Marriage(body.Marriage)).
		SetMedicalHistory(body.MedicalHistory).
		SetMedicationStatus(body.MedicationStatus).
		SetDementedAmongDirectRelatives(body.DementedAmongDirectRelatives).
		SetHeadInjuryExperience(body.HeadInjuryExperience).
		SetEarCondition(user.EarCondition(body.EarCondition)).
		SetEyesightCondition(user.EyesightCondition(body.EyesightCondition)).
		SetSmokingHabit(user.SmokingHabit(body.SmokingHabit)).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userNode)
}

// PUT		/users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var body UserBody
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

	userNode, err := h.DB.User.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedUserNode, err := h.DB.User.
		UpdateOne(userNode).
		SetBirthYear(body.BirthYear).
		SetHeight(body.Height).
		SetWeight(body.Weight).
		SetGender(user.Gender(body.Gender)).
		SetEducationLevel(user.EducationLevel(body.EducationLevel)).
		SetOccupation(user.Occupation(body.Occupation)).
		SetMarriage(user.Marriage(body.Marriage)).
		SetMedicalHistory(body.MedicalHistory).
		SetMedicationStatus(body.MedicationStatus).
		SetDementedAmongDirectRelatives(body.DementedAmongDirectRelatives).
		SetHeadInjuryExperience(body.HeadInjuryExperience).
		SetEarCondition(user.EarCondition(body.EarCondition)).
		SetEyesightCondition(user.EyesightCondition(body.EyesightCondition)).
		SetSmokingHabit(user.SmokingHabit(body.SmokingHabit)).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUserNode)
}

// DELETE 	/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	if err := h.DB.User.DeleteOneID(c.Param("id")).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
