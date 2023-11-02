package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

//	@Summary				Get Own User
//	@Description.markdown	user_self.get
//	@Tags					User
//	@Produce				json
//	@Success				200	{object}	[]ent.User
//	@Router					/user [get]
func (h *Handler) GetSelf(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	userNode, err := h.GetUserById(c.Request.Context(), *userId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, userNode)
}

//	@Summary				Create HealthKit Data
//	@Description.markdown	user_healthkit.post
//	@Tags					User
//	@Accept					json
//	@Produce				json
//	@Param					healthkit	body		types.BaseHealthKit	true	"The healthkit to be created"
//	@Success				200			{object}	ent.HealthKit
//	@Router					/user/healthkit [post]
func (h *Handler) CreateUserHealthKitData(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var body types.BaseHealthKit
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

	healthkitNode, err := h.DB.HealthKit.
		Create().
		SetUserID(*userId).
		SetStartTime(body.StartTime).
		SetEndTime(body.EndTime).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, datum := range body.Data {
		if len(datum) != 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "the data should have 6 columns"})
			return
		}

		_, err := h.DB.HKData.
			Create().
			SetHealthkit(healthkitNode).
			SetType(datum[0]).
			SetValue(datum[1]).
			SetDataID(datum[2]).
			SetStartTimestamp(datum[3]).
			SetEndTimestamp(datum[4]).
			SetTimezoneID(datum[5]).
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, healthkitNode)
}

//	@Summary				Update User's FCM Token
//	@Description.markdown	user_fcm.put
//	@Tags					User
//	@Accept					json
//	@Produce				json
//	@Param					healthkit	body		types.FcmTokenRequest	true	"The FCM token to update"
//	@Success				200			{object}	ent.User
//	@Router					/user/fcm [put]
func (h *Handler) UpdateUserFcmToken(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var body types.FcmTokenRequest
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

	userNode, err := h.DB.User.UpdateOneID(*userId).
		SetFcmToken(body.FcmToken).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userNode)
}

//	@Summary				Delete User
//	@Description.markdown	user.delete
//	@Tags					User
//	@Success				204
//	@Router					/user [delete]
func (h *Handler) DeleteSelf(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if err := h.DB.User.DeleteOneID(*userId).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
