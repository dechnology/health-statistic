package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope       string   `json:"scope"`
	Permissions []string `json:"permissions"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// CheckScope checks whether our claims have a specific scope.
func (c CustomClaims) CheckScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

func Authenticate() gin.HandlerFunc {
	issuerUrl, err := url.Parse(os.Getenv("AUTH0_ISSUER_URL"))
	if err != nil {
		log.Fatalf("failed to parse Auth0 issuer url: %v", err)
	}

	audience := os.Getenv("AUTH0_AUDIENCE")

	log.Printf("issuer: %v\naudience: %v", issuerUrl.String(), audience)

	provider := jwks.NewCachingProvider(issuerUrl, 5*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerUrl.String(),
		[]string{audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	// return func(c *gin.Context) {
	// 	tokenString := strings.TrimPrefix(
	// 		c.GetHeader("Authorization"), "Bearer ",
	// 	)
	// 	log.Printf("token receieved: %v", tokenString)

	// 	_, err := jwtValidator.ValidateToken(
	// 		c.Request.Context(),
	// 		tokenString,
	// 	)
	// 	if err != nil {
	// 		log.Printf("Encountered error while validating JWT: %v", err)
	// 		c.JSON(
	// 			http.StatusUnauthorized,
	// 			gin.H{"error": "Failed to validate JWT."},
	// 		)
	// 		c.Abort()
	// 		return
	// 	}

	// 	c.Next()
	// }

	// Original non-gin's way
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return adapter.Wrap(middleware.CheckJWT)
}

func Authorize(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Context().
			Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		if !claims.CheckScope(scope) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": fmt.Sprintf("missing scope: %v", scope),
			})
			c.Abort() // Abort the request, as the user doesn't have the required scope
			return
		}
		c.Next()
	}
}
