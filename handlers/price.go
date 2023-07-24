package handlers

import (
	"net/http"

	"github.com/eesoymilk/health-statistic-api/ent/price"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary				Get Prices
//	@Description.markdown	prices.get
//	@Tags					Price
//	@Produce				json
//	@Success				200	{object}	[]ent.Price
//	@Router					/prices [get]
func (h *PriceHandler) GetPrices(c *gin.Context) {
	prices, err := h.DB.Price.
		Query().
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prices)
}

//	@Summary				Get Price
//	@Description.markdown	price.get
//	@Tags					Price
//	@Produce				json
//	@Param					id	path		string	true	"The price's ID"
//	@Success				200	{object}	ent.Price
//	@Router					/prices/{id} [get]
func (h *PriceHandler) GetPrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prices, err := h.DB.Price.
		Query().
		Where(price.ID((id))).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, prices)
}

//	@Summary				Delete Price
//	@Description.markdown	price.delete
//	@Tags					Price
//	@Produce				json
//	@Param					id	path	string	true	"The price's ID."
//	@Success				200
//	@Router					/prices/{id} [delete]
func (h *PriceHandler) DeletePrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Price.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
