package router

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

/********************************  TODO  ********************************/
/*	1. GET  	/users													*/
/*		- fetch all users 												*/
/*	2. POST 	/users													*/
/*		- create a new user 											*/
/*	3. GET  	/users/{id}												*/
/*		- fetch a specific user 										*/
/*	4. PUT  	/users/{id}												*/
/*		- update a specific user 										*/
/*	5. DELETE	/users/{id}												*/
/*		- delete a specific user 										*/
/* ====================================================================	*/
/*	6. GET  	/questionnaires											*/
/*		- fetch all questionnaires 										*/
/*	7. POST 	/questionnaires											*/
/*		- create a new questionnaire 									*/
/*	8. GET  	/questionnaires/{id}									*/
/*		- fetch a specific questionnaire and its reponses				*/
/*  9. POST		/questionnaires/{id}									*/
/*		- create a questionnaire response								*/
/* 10. DELETE	/questionnaires/{id}									*/
/*		- delete a specific questionnaire 								*/
/* ====================================================================	*/
/* 11. GET		/responses												*/
/*		- fetch all responses											*/
/* 12. GET		/responses/{id}											*/
/*		- fetch a specific response										*/
/* 13. DELETE	/responses/{id}											*/
/*		- delete a specific response									*/
/************************************************************************/

// New registers the routes and returns the router.
func New(db *ent.Client) *gin.Engine {
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

	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			h := handlers.UserHandler{DB: db}
			userGroup.GET("/", h.GetAllUsers)
			userGroup.GET("/:id", h.GetUserById)
			userGroup.POST("/", h.CreateUser)
			userGroup.PUT("/:id", h.UpdateUser)
			userGroup.DELETE("/:id", h.DeleteUserById)
		}

		questionnaireGroup := v1.Group("/questionnaires")
		{
			h := handlers.QuestionnaireHandler{DB: db}
			questionnaireGroup.GET("/", h.GetAllQuestionnaires)
			questionnaireGroup.GET("/:id", h.GetQuestionnaireById)
			questionnaireGroup.POST("/", h.CreateQuestionnaire)
			questionnaireGroup.POST("/:id", h.CreateResponse)
			questionnaireGroup.DELETE("/:id", h.DeleteQuestionnaireById)

		}

	}

	return r
}

func fecthAllUsers(c *gin.Context) {

}
