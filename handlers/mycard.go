package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/gin-gonic/gin"
)

//	@Summary				Get MyCards
//	@Description.markdown	mycards.get
//	@Tags					MyCard
//	@Produce				json
//	@Success				200	{object}	[]ent.MyCard
//	@Router					/mycards [get]
func (h *MyCardHandler) GetMyCards(c *gin.Context) {
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
func (h *MyCardHandler) GetMyCard(c *gin.Context) {
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

//	@Summary				Delete MyCard
//	@Description.markdown	mycard.delete
//	@Tags					MyCard
//	@Produce				json
//	@Param					id	path	string	true	"The mycard's ID."
//	@Success				200
//	@Router					/mycards/{id} [delete]
func (h *MyCardHandler) DeleteMyCard(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.MyCard.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
