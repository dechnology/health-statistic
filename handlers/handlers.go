package handlers

import (
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

// Handler for /api/public endpoint
func HandlePublic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from a public endpoint! You don't need to be authenticated to see this.",
	})
}

// Handler for /api/private endpoint
func HandlePrivate(c *gin.Context) {
	claims, ok := c.Request.Context().Value(
		jwtmiddleware.ContextKey{},
	).(*validator.ValidatedClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to get validated JWT claims."},
		)
		return
	}

	customClaims, ok := claims.CustomClaims.(*middlewares.CustomClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to cast custom JWT claims to specific type."},
		)
		return
	}

	log.Print(*customClaims)

	if len(customClaims.Username) == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]string{"message": "Username in JWT claims was empty."},
		)
		return
	}

	c.JSON(http.StatusOK, claims)
}

// TODO: Handler for /api/private-scoped endpoint
func HandlePrivateScoped(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello from a private endpoint! You need to be authenticated to see this.",
		},
	)
}
