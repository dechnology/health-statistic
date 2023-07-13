package middlewares

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Scope    string `json:"scope"`
	// ShouldReject bool   `json:"shouldReject,omitempty"`
}

// Validate errors out if `ShouldReject` is true.
func (c *CustomClaims) Validate(ctx context.Context) error {
	// if c.ShouldReject {
	// 	return errors.New("should reject was set to true")
	// }
	return nil
}

// EnsureValidToken is a gin.HandlerFunc middleware
// that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	issuerURL, err := url.Parse(os.Getenv("AUTH0_ISSUER_URL"))
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return adapter.Wrap(middleware.CheckJWT)

	// return func(ctx *gin.Context) {
	// 	encounteredError := true
	// 	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 		encounteredError = false
	// 		ctx.Request = r
	// 		ctx.Next()
	// 	}

	// 	middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

	// 	if encounteredError {
	// 		ctx.AbortWithStatusJSON(
	// 			http.StatusUnauthorized,
	// 			map[string]string{"message": "JWT is invalid."},
	// 		)
	// 	}
	// }
}
