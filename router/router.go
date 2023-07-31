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
	{
		// Health check endpoint for ELB
		// This endpoint does not require an Auth0 token
		v1.GET("/health_check", handlers.HealthCheck)

		// All routes below require an Auth0 token
		v1.Use(middlewares.Authenticate())

		register := v1.Group("/register")
		{
			register.POST("", h.Register)
		}

		user := v1.Group("/user")
		{
			user.GET("", h.GetSelf)
			user.PUT("")
			user.DELETE("")
			user.GET("/notifications")
		}

		readUsers := middlewares.Authorize("read:users")
		writeUsers := middlewares.Authorize("write:users")
		users := v1.Group("/users")
		users.Use(readUsers)
		{
			users.GET("", h.GetUsers)

			id := users.Group("/:id")
			{
				id.GET("", h.GetUser)
				id.PUT("", writeUsers, h.UpdateUser)
				id.DELETE("", writeUsers, h.DeleteUser)
				id.GET("/notifications", h.GetUserNotifications)
			}
		}

		// readQuestionnaires := middlewares.Authorize("read:questionnaires")
		writeQuestionnaires := middlewares.Authorize("write:questionnaires")
		questionnaires := v1.Group("/questionnaires")
		// questionnaires.Use(readQuestionnaires)
		{
			questionnaires.GET("", h.GetQuestionnaires)
			questionnaires.GET("/registration", h.GetRegistrationQuestionnaire)
			questionnaires.POST("", writeQuestionnaires, h.CreateQuestionnaire)

			id := questionnaires.Group("/:id")
			{
				id.GET("", h.GetQuestionnaire)
				id.DELETE("", writeQuestionnaires, h.DeleteQuestionnaire)
				id.POST("/responses", writeQuestionnaires, h.CreateQuestionnaireResponse)
				id.GET("/responses", h.GetQuestionnaireResponses)
			}
		}

		readMyCards := middlewares.Authorize("read:mycards")
		writeMyCards := middlewares.Authorize("write:mycards")
		mycards := v1.Group("/mycards")
		mycards.Use(readMyCards)
		{
			mycards.GET("", h.GetMyCards)
			mycards.GET("/:id", h.GetMyCard)
			mycards.DELETE("/:id", writeMyCards, h.DeleteMyCard)
		}

		// readPrices := middlewares.Authorize("read:prices")
		writePrices := middlewares.Authorize("write:prices")
		prices := v1.Group("/prices")
		// prices.Use(readPrices)
		{
			prices.GET("", h.GetPrices)
			prices.GET("/:id", h.GetPrice)
			prices.DELETE("/:id", writePrices, h.DeletePrice)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
