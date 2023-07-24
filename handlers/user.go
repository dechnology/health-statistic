package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/price"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

//	@Summary				Get Users
//	@Description.markdown	users.get
//	@Tags					User
//	@Produce				json
//	@Success				200	{object}	[]ent.User
//	@Router					/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.DB.User.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

//	@Summary				Get User
//	@Description.markdown	user.get
//	@Tags					User
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	ent.User
//	@Router					/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	user, err := h.DB.User.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

//	@Summary				Create User
//	@Description.markdown	user.post
//	@Tags					User
//	@Accept					json
//	@Produce				json
//	@Param					user	body		types.BaseUser	true	"The user to be created"
//	@Success				200		{object}	ent.User
//	@Router					/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
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

//	@Summary				Update User
//	@Description.markdown	user.put
//	@Tags					User
//	@Accept					json
//	@Produce				json
//	@Param					id		path		string			true	"The user's Auth0 ID"
//	@Param					user	body		types.BaseUser	true	"user to be updated"
//	@Success				200		{object}	ent.User
//	@Router					/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
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

//	@Summary				Delete User
//	@Description.markdown	user.delete
//	@Tags					User
//	@Produce				json
//	@Param					id	path	string	true	"The user's Auth0 ID"
//	@Success				200
//	@Router					/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	if err := h.DB.User.DeleteOneID(c.Param("id")).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary				Get All Notifications From an User
//
//	@Description.markdown	user_notifications.get
//
//	@Tags					User
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	[]ent.Notifications
//
//	@Router					/users/{id}/notifications [get]
func (h *UserHandler) GetUserNotifications(c *gin.Context) {
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

//	@Summary				Get All MyCards From an User
//	@Description.markdown	user_mycards.get
//	@Tags					User
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	[]ent.MyCard
//	@Router					/users/{id}/mycards [get]
func (h *UserHandler) GetUserMyCards(c *gin.Context) {
	mycards, err := h.DB.MyCard.
		Query().
		Where(
			mycard.HasRecipientWith(user.ID(c.Param("id"))),
		).
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mycards)
}

//	@Summary				Get All Prices From an User
//	@Description.markdown	user_prices.get
//	@Tags					User
//	@Produce				json
//	@Param					id	path		string	true	"The user's Auth0 ID"
//	@Success				200	{object}	[]ent.Price
//	@Router					/users/{id}/prices [get]
func (h *UserHandler) GetUserPrices(c *gin.Context) {
	prices, err := h.DB.Price.
		Query().
		Where(
			price.HasRecipientWith(user.ID(c.Param("id"))),
		).
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prices)
}
