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
//	@Tags					Self
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

//	@Summary				Create HealthKit Datum
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
		SetStartDate(body.StartDate).
		SetEndDate(body.EndDate).
		SetStepCount(body.StepCount).
		SetUserID(*userId).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, healthkitNode)
}
