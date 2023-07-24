package router

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
/* ======================================================================== */
/* 16. GET		/mycards													*/
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
		// Health check endpoint for ELB
		v1.GET("/health_check", handlers.HealthCheck)

		userGroup := v1.Group("/users")
		{
			h := handlers.UserHandler{DB: db}
			userGroup.GET("/", h.GetUsers)
			userGroup.POST("/", h.CreateUser)

			idGroup := userGroup.Group("/:id")
			{
				idGroup.GET("/", h.GetUser)
				idGroup.PUT("/", h.UpdateUser)
				idGroup.DELETE("/", h.DeleteUser)
				idGroup.GET("/notifications", h.GetUserNotifications)
				idGroup.GET("/mycards", h.GetUserMyCards)
				idGroup.GET("/prices", h.GetUserPrices)
			}
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

		notificationGroup := v1.Group("/notifications")
		{
			h := handlers.NotificationHandler{DB: db}
			notificationGroup.GET("/", h.GetNotifications)
			notificationGroup.GET("/:id", h.GetNotification)
			notificationGroup.DELETE("/:id", h.DeleteNotification)
		}

		myCardGroup := v1.Group("/mycards")
		{
			h := handlers.MyCardHandler{DB: db}
			myCardGroup.GET("/", h.GetMyCards)
			myCardGroup.GET("/:id", h.GetMyCard)
			myCardGroup.DELETE("/:id", h.DeleteMyCard)
		}

		priceGroup := v1.Group("/prices")
		{
			h := handlers.PriceHandler{DB: db}
			priceGroup.GET("/", h.GetPrices)
			priceGroup.GET("/:id", h.GetPrice)
			priceGroup.DELETE("/:id", h.DeletePrice)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
