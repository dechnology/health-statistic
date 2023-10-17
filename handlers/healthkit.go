package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//	@Summary				Get HealthKitData
//	@Description.markdown	healthkit.get
//	@Tags					HealthKit
//	@Produce				json
//	@Success				200	{object}	map[string]interface{}
//	@Router					/healthkit [get]
func (h *Handler) GetHealthKitData(c *gin.Context) {
	healthkitData, err := h.DB.HealthKit.
		Query().
		All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, healthkitData)
}

//	@Summary				Create HealthKitDatum
//	@Description.markdown	healthkit.post
//	@Tags					HealthKit
//	@Accept					json
//	@Produce				json
//	@Param					healthkit	body		map[string]interface{}	true	"The healthkit to be created"
//	@Success				200			{object}	map[string]interface{}
//	@Router					/healthkit [post]
func (h *Handler) CreateHealthKitData(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user_id exists in the JSON body
	userID, exists := body["user_id"]
	if !exists || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
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
		SetData(body).
		SetUserID(userID.(string)). // Setting the userID
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, healthkitNode)
}
