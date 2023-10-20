package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

//	@Summary				Get Users
//	@Description.markdown	users.get
//	@Tags					Users
//	@Produce				json
//	@Success				200	{object}	[]ent.User
//	@Router					/users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.DB.User.Query().
		WithMycards().
		WithPrices().
		WithHealthkit().
		All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

//	@Summary				Get User
//	@Description.markdown	user.get
//	@Tags					Users
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	ent.User
//	@Router					/users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	userNode, err := h.GetUserById(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, userNode)
}

//	@Summary				Update User
//	@Description.markdown	user.put
//	@Tags					Users
//	@Accept					json
//	@Produce				json
//	@Param					id		path		string			true	"The user's Auth0 ID"
//	@Param					user	body		types.BaseUser	true	"user to be updated"
//	@Success				200		{object}	ent.User
//	@Router					/users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var body types.BaseUser
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
		SetMedicalHistory(user.MedicalHistory(body.MedicalHistory)).
		SetMedicationStatus(user.MedicationStatus(body.MedicationStatus)).
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

//	@Summary				Delete User
//	@Description.markdown	user.delete
//	@Tags					Users
//	@Produce				json
//	@Param					id	path	string	true	"The user's Auth0 ID"
//	@Success				200
//	@Router					/users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	if err := h.DB.User.DeleteOneID(c.Param("id")).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary				Get All Notifications From an User
//	@Description.markdown	user_notifications.get
//	@Tags					Users
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	[]ent.Notifications
//
//	@Router					/users/{id}/notifications [get]
func (h *Handler) GetUserNotifications(c *gin.Context) {
	notifications, err := h.DB.Notification.
		Query().
		Where(notification.HasRecipientWith(user.ID(c.Param("id")))).
		All(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, notifications)
}
