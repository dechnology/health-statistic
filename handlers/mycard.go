package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

//	@Summary				Get MyCards
//	@Description.markdown	mycards.get
//	@Tags					MyCard
//	@Produce				json
//	@Success				200	{object}	[]ent.MyCard
//	@Router					/mycards [get]
func (h *Handler) GetMyCards(c *gin.Context) {
	mycards, err := h.DB.MyCard.
		Query().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mycards)
}

//	@Summary				Get MyCard
//	@Description.markdown	mycard.get
//	@Tags					MyCard
//	@Produce				json
//	@Param					id	path		string	true	"The mycard's ID"
//	@Success				200	{object}	ent.MyCard
//	@Router					/mycards/{id} [get]
func (h *Handler) GetMyCard(c *gin.Context) {
	id := c.Param("id")

	mycards, err := h.DB.MyCard.
		Query().
		Where(mycard.ID((id))).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, mycards)
}

//	@Summary				Create MyCard
//	@Description.markdown	mycard.post
//	@Tags					MyCard
//	@Accept					json
//	@Produce				json
//	@Param					mycard	body		types.BaseMyCard	true	"The mycard to be created"
//	@Success				200		{object}	ent.MyCard
//	@Router					/mycards [post]
func (h *Handler) CreateMyCard(c *gin.Context) {
	var body types.BaseMyCard
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

	mycardNode, err := h.DB.MyCard.
		Create().
		SetID(body.CardNumber).
		SetCardPassword(body.CardPassword).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mycardNode)
}

//	@Summary				Delete MyCard
//	@Description.markdown	mycard.delete
//	@Tags					MyCard
//	@Produce				json
//	@Param					id	path	string	true	"The mycard's ID."
//	@Success				200
//	@Router					/mycards/{id} [delete]
func (h *Handler) DeleteMyCard(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.MyCard.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
