package handlers

import (
	"net/http"

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
