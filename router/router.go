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
	r.Use(middlewares.CorsMiddleware())

	h := handlers.Handler{DB: db}

	v1 := r.Group("/api/v1")
	{
		// Health check endpoint for ELB
		v1.GET("/health_check", handlers.HealthCheck)

		register := v1.Group("/register")
		{
			register.POST("", h.Register)
		}

		users := v1.Group("/users")
		{
			users.GET("", h.GetUsers)

			id := users.Group("/:id")
			{
				id.GET("", h.GetUser)
				id.PUT("", h.UpdateUser)
				id.DELETE("", h.DeleteUser)
				id.GET("/notifications", h.GetUserNotifications)
				id.GET("/mycards", h.GetUserMyCards)
				id.GET("/prices", h.GetUserPrices)
			}
		}

		questionnaires := v1.Group("/questionnaires")
		{
			questionnaires.GET("", h.GetQuestionnaires)
			questionnaires.POST("", h.CreateQuestionnaire)
			questionnaires.GET(
				"/registration",
				h.GetRegistrationQuestionnaire,
			)

			id := questionnaires.Group("/:id")
			{
				id.GET("", h.GetQuestionnaire)
				id.DELETE("", h.DeleteQuestionnaire)
				id.POST("/new/question", h.CreateQuestions)
				id.POST("/new/response", h.CreateQuestionnaireResponse)
				id.GET("/responses", h.GetQuestionnaireResponses)
			}
		}

		responses := v1.Group("/responses")
		{
			responses.GET("", h.GetResponses)
			responses.GET("/:id", h.GetResponse)
			responses.DELETE("/:id", h.DeleteResponse)
		}

		questions := v1.Group("/questions")
		{
			questions.GET("", h.GetQuestions)
			questions.GET("/:id", h.GetQuestion)
			questions.DELETE("/:id", h.DeleteQuestion)
		}

		notifications := v1.Group("/notifications")
		{
			notifications.GET("", h.GetNotifications)
			notifications.GET("/:id", h.GetNotification)
			notifications.DELETE("/:id", h.DeleteNotification)
		}

		mycards := v1.Group("/mycards")
		{
			mycards.GET("", h.GetMyCards)
			mycards.GET("/:id", h.GetMyCard)
			mycards.DELETE("/:id", h.DeleteMyCard)
		}

		prices := v1.Group("/prices")
		{
			prices.GET("", h.GetPrices)
			prices.GET("/:id", h.GetPrice)
			prices.DELETE("/:id", h.DeletePrice)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
