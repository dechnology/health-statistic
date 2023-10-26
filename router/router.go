package router

import (
	"firebase.google.com/go/v4/messaging"
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(db *ent.Client, fcm *messaging.Client) *gin.Engine {
	r := gin.Default()

	h := handlers.Handler{
		DB:  db,
		FCM: fcm,
		RegistrationQuestionnaireId: uuid.MustParse(
			"88888888-8888-4888-8888-888888888888",
		),
	}

	// adminAuth := middlewares.Authorize("admin")
	r.Use(middlewares.CorsMiddleware())

	v1 := r.Group("/api/v1")
	{
		// Health check endpoint for ELB
		// This endpoint does not require an Auth0 token
		v1.GET("/health_check", h.HealthCheck)

		// All routes below require an Auth0 token
		v1.Use(middlewares.Authenticate())

		register := v1.Group("/register")
		{
			register.POST("", h.Register)
		}

		healthkit := v1.Group("/healthkit")
		{
			healthkit.GET("", h.GetHealthKitData)
		}

		user := v1.Group("/user")
		{
			user.GET("", h.GetSelf)
			user.POST("/healthkit", h.CreateUserHealthKitData)

			// TODO
			// user.PUT("")
			// user.DELETE("")
			// user.GET("/notifications")
		}

		deegoo := v1.Group("/deegoo")
		// deegoo.Use(adminAuth)
		{
			deegoo.POST("", h.SubmitDeegoo)
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

		writePrices := middlewares.Authorize("write:prices")
		prices := v1.Group("/prices")
		{
			prices.GET("", h.GetPrices)
			prices.GET("/:id", h.GetPrice)
			prices.DELETE("/:id", writePrices, h.DeletePrice)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
