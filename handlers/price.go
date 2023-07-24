package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent/price"
	"github.com/eesoymilk/health-statistic-api/types"
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

//	@Summary				Create Price
//	@Description.markdown	price.post
//	@Tags					Price
//	@Accept					json
//	@Produce				json
//	@Param					price	body		types.BasePrice	true	"The price to be created"
//	@Success				200		{object}	ent.Price
//	@Router					/prices [post]
func (h *PriceHandler) CreatePrice(c *gin.Context) {
	var body types.BasePrice
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

	priceNode, err := h.DB.Price.
		Create().
		SetName(body.Name).
		SetDescription(body.Description).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, priceNode)
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
