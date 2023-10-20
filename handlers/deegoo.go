package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	c.JSON(http.StatusOK, deegooNode)
}
