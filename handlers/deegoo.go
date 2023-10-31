package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/v4/messaging"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

// @Summary				Create Deegoo
// @Description.markdown	deegoo.post
// @Tags					Deegoo
// @Accept					json
// @Produce				json
// @Param					deegoo	body		types.BaseDeegoo	true	"The deegoo scores to submit"
// @Success				200		{object}	ent.Deegoo
// @Router					/deegoo [post]
func (h *Handler) SubmitDeegoo(c *gin.Context) {
	var body types.BaseDeegoo
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

	deegooNode, err := h.DB.Deegoo.
		Create().
		SetPerception(body.Perception).
		SetFocus(body.Focus).
		SetExecution(body.Execution).
		SetMemory(body.Memory).
		SetLanguage(body.Language).
		SetUserID(body.UserId).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "dTHzcCQDCUZcmnU6kX7u5f:APA91bHkdnCCqGtMVKN9XX4k9hz58tyUi79PlF9LEF5D6LkydkUJmpfWpMjDb6Y1Y5ruHlppuHFLGwIC9jjx3rkzOXtOkC2eR7yRzSDRUO6iXnIcG_GdaN46zjTJrBl5lgWBCxbc7u_r"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := h.FCM.Send(c.Request.Context(), message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	c.JSON(http.StatusOK, deegooNode)
}
