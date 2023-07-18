package router

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

/**********************************  TODO  **********************************/
/*	1. GET  	/users														*/
/*		- fetch all users 													*/
/*	2. POST 	/users														*/
/*		- create a new user 												*/
/*	3. GET  	/users/:id													*/
/*		- fetch a specific user 											*/
/*	4. PUT  	/users/:id													*/
/*		- update a specific user 											*/
/*	5. DELETE	/users/:id													*/
/*		- delete a specific user 											*/
/* ======================================================================== */
/*	6. GET  	/questionnaires												*/
/*		- fetch all questionnaires with its questions and responses 		*/
/*	7. POST 	/questionnaires												*/
/*		- create a new empty questionnaire									*/
/*	8. GET  	/questionnaires/:id											*/
/*		- fetch a specific questionnaire  with its questions and responses	*/
/*  9. DELETE	/questionnaires/:id											*/
/*		- delete a specific questionnaire 									*/
/* 10. POST 	/questionnaires/:id/new/question							*/
/*		- create a new question for a specific questionnaire 				*/
/* 11. POST 	/questionnaires/:id/new/response							*/
/*		- create a new response for a specific questionnaire 				*/
/* ======================================================================== */
/* 12. GET		/responses													*/
/*		- fetch all responses												*/
/* 14. GET		/responses/:id												*/
/*		- fetch a specific response											*/
/* 15. DELETE	/responses/:id												*/
/*		- delete a specific response										*/
/* ======================================================================== */
/* 16. GET		/questions													*/
/*		- fetch all questions												*/
/* 18. GET		/questions/:id												*/
/*		- fetch a specific question											*/
/* 19. DELETE	/questions/:id												*/
/*		- delete a specific question										*/
/****************************************************************************/

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
			userGroup.GET("/", h.GetUsers)
			userGroup.POST("/", h.CreateUser)
			userGroup.GET("/:id", h.GetUser)
			userGroup.PUT("/:id", h.UpdateUser)
			userGroup.DELETE("/:id", h.DeleteUser)
		}

		questionnaireGroup := v1.Group("/questionnaires")
		{
			h := handlers.QuestionnaireHandler{DB: db}
			questionnaireGroup.GET("/", h.GetQuestionnaires)
			questionnaireGroup.POST("/", h.CreateQuestionnaire)

			idGroup := questionnaireGroup.Group("/:id")
			{
				idGroup.GET("/", h.GetQuestionnaire)
				idGroup.DELETE("/", h.DeleteQuestionnaire)
				idGroup.POST("/new/question", h.CreateQuestion)
				idGroup.POST("/new/response", h.CreateResponse)
			}
		}

		responseGroup := v1.Group("/responses")
		{
			h := handlers.ResponseHandler{DB: db}
			responseGroup.GET("/", h.GetResponses)
			responseGroup.GET("/:id", h.GetResponse)
			responseGroup.DELETE("/:id", h.DeleteResponse)
		}

		questionGroup := v1.Group("/questions")
		{
			h := handlers.QuestionHandler{DB: db}
			questionGroup.GET("/", h.GetQuestions)
			questionGroup.GET("/:id", h.GetQuestion)
			questionGroup.DELETE("/:id", h.DeleteQuestion)
		}
	}

	return r
}
