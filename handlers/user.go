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

// @Summary     Get Users
// @Description Get all users from the database
// @Tags        User
// @Produce     json
// @Success 	200 {object} []ent.User
// @Router      /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.DB.User.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary     Get User
// @Description Get a user by an Auth0 ID.
// @Tags        User
// @Produce     json
// @Param		id path string true "The user's Auth0 ID"
// @Success 	200 {object} ent.User
// @Router      /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	user, err := h.DB.User.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary     Create User
// @Description Create a new user
// @Tags        User
// @Accept		json
// @Produce     json
// @Param		user body ent.User true "The user to be created"
// @Success 	200 {object} ent.User
// @Router      /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var body ent.User
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

// @Summary     Update User
// @Description update a user with specified Auth0 ID
// @Tags        User
// @Accept		json
// @Produce     json
// @Param		id path string true "The user's Auth0 ID"
// @Param		user body ent.User true "user to be updated"
// @Success 	200 {object} ent.User
// @Router      /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var body ent.User
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

// @Summary     Delete User
// @Description Delete a user by Auth0 ID
// @Tags        User
// @Produce     json
// @Param		id path string true "The user's Auth0 ID"
// @Success 	200
// @Router      /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	if err := h.DB.User.DeleteOneID(c.Param("id")).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
