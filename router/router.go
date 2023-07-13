package router

import (
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	r := gin.Default()

	// This route is always accessible.
	r.GET(
		"/api/public",
		handlers.HandlePublic,
	)

	// This route is only accessible if the user has a valid access_token.
	r.GET(
		"/api/private",
		middlewares.EnsureValidToken(),
		handlers.HandlePrivate,
	)

	// This route is only accessible if the user has a
	// valid access_token with the read:messages scope.
	r.GET(
		"/api/private-scoped",
		middlewares.EnsureValidToken(),
		handlers.HandlePrivateScoped,
	)

	return r
}
