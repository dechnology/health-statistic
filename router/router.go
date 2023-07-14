package router

import (
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

/********************  TODO  ********************/
/*	1. GET  	/users							*/
/*		- fetch all users 						*/
/*	2. GET  	/users/{id}						*/
/*		- fetch a specific user 				*/
/*	3. POST 	/users							*/
/*		- create a new user 					*/
/*	4. PUT  	/users/{id}						*/
/*		- update a specific user 				*/
/*	5. DELETE	/users/{id}						*/
/*		- delete a specific user 				*/
/* ============================================ */
/*	6. GET  	/questionnaires					*/
/*		- fetch all questionnaires 				*/
/*	7. GET  	/questionnaires/{id}			*/
/*		- fetch a specific questionnaire		*/
/*	8. POST 	/questionnaires					*/
/*		- create a new questionnaire 			*/
/*	9. PUT  	/questionnaires/{id}			*/
/*		- update a specific questionnaire 		*/
/* 10. DELETE	/questionnaires/{id}			*/
/*		- delete a specific questionnaire 		*/
/* ============================================ */
/* 11. GET		/questionnaires/responses		*/
/*		- fetch all responses					*/
/* 12. GET		/questionnaires/responses/{id}	*/
/*		- fetch a specific response				*/
/* 13. POST		/questionnaires/responses		*/
/*		- create a questionnaire response 		*/
/* 14. DELETE	/questionnaires/responses/{id}	*/
/*		- delete a specific response 			*/
/************************************************/

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

	userGroup := r.Group("api/users")
	{
		userGroup.GET("/")
		userGroup.GET("/:id")
		userGroup.POST("/")
		userGroup.PUT("/:id")
		userGroup.DELETE("/:id")
	}

	return r
}
