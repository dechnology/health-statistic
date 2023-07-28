package router

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(db *ent.Client) *gin.Engine {
	r := gin.Default()

	h := handlers.Handler{
		DB: db,
		RegistrationQuestionnaireId: uuid.MustParse(
			"88888888-8888-4888-8888-888888888888",
		),
	}

	r.Use(middlewares.CorsMiddleware())

	v1 := r.Group("/api/v1")
	// v1.Use(middlewares.Authenticate())
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
				id.POST("/responses", h.CreateQuestionnaireResponse)
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
