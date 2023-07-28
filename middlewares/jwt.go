package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope       string   `json:"scope"`
	Permissions []string `json:"permissions"`
}

type Token struct {
	CustomClaims     *CustomClaims
	RegisteredClaims *validator.RegisteredClaims
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

func GetTokenInGin(c *gin.Context) (*Token, error) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token not found in context.",
		})
		c.Abort() // Abort the request, as the token was not found
		return nil, errors.New("token not found in context")
	}
	validatedClaims := token.(*validator.ValidatedClaims)
	customClaims := validatedClaims.CustomClaims.(*CustomClaims)

	return &Token{
		RegisteredClaims: &validatedClaims.RegisteredClaims,
		CustomClaims:     customClaims,
	}, nil

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

	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(
			c.GetHeader("Authorization"), "Bearer ",
		)
		log.Printf("token receieved: %v", tokenString)

		validatedClaims, err := jwtValidator.ValidateToken(
			c.Request.Context(),
			tokenString,
		)
		if err != nil {
			log.Printf("Encountered error while validating JWT: %v", err)
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "Failed to validate JWT."},
			)
			c.Abort()
			return
		}

		c.Set("token", validatedClaims.(*validator.ValidatedClaims))
		c.Next()
	}
}

func Authorize(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetTokenInGin(c)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": err.Error(),
				},
			)
			c.Abort()
			return
		}

		// Check for the provided scope
		if !token.CustomClaims.CheckScope(scope) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": fmt.Sprintf("missing scope: %v", scope),
			})
			c.Abort() // Abort the request, as the user doesn't have the required scope
			return
		}

		c.Next()
	}
}
