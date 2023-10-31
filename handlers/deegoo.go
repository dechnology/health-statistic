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

//	@Summary				Create Deegoo
//	@Description.markdown	deegoo.post
//	@Tags					Deegoo
//	@Accept					json
//	@Produce				json
//	@Param					deegoo	body		types.BaseDeegoo	true	"The deegoo scores to submit"
//	@Success				200		{object}	ent.Deegoo
//	@Router					/deegoo [post]
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

	user, err := h.GetUserById(c.Request.Context(), body.UserId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deegooNode, err := h.DB.Deegoo.
		Create().
		SetPerception(body.Perception).
		SetFocus(body.Focus).
		SetExecution(body.Execution).
		SetMemory(body.Memory).
		SetLanguage(body.Language).
		SetUser(user).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.FcmToken == "" {
		c.JSON(http.StatusOK, deegooNode)
		return
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Messgae from DeeGoo",
			Body:  "Your DeeGoo score has been submitted!",
		},
		Token: user.FcmToken,
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
