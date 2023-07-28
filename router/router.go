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
	v1.Use(middlewares.Authenticate())
	{
		// Health check endpoint for ELB
		v1.GET("/health_check", handlers.HealthCheck)

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

		users := v1.Group("/users")
		{
			users.GET(
				"",
				middlewares.Authorize("read:users"),
				h.GetUsers,
			)

			id := users.Group("/:id")
			{
				id.GET(
					"",
					middlewares.Authorize("read:current_user"),
					h.GetUser,
				)
				id.PUT(
					"",
					middlewares.Authorize("write:current_user"),
					h.UpdateUser,
				)
				id.DELETE(
					"",
					middlewares.Authorize("write:current_user"),
					h.DeleteUser,
				)
				id.GET(
					"/notifications",
					middlewares.Authorize("read:current_user"),
					h.GetUserNotifications,
				)
			}
		}

		questionnaires := v1.Group("/questionnaires")
		{
			questionnaires.GET(
				"",
				middlewares.Authorize("read:questionnaires"),
				h.GetQuestionnaires,
			)
			questionnaires.GET(
				"/registration",
				middlewares.Authorize("read:questionnaires"),
				h.GetRegistrationQuestionnaire,
			)
			questionnaires.POST(
				"",
				middlewares.Authorize("write:questionnaires"),
				h.CreateQuestionnaire,
			)

			id := questionnaires.Group("/:id")
			{
				id.GET(
					"",
					middlewares.Authorize("read:questionnaires"),
					h.GetQuestionnaire,
				)
				id.DELETE(
					"",
					middlewares.Authorize("write:questionnaires"),
					h.DeleteQuestionnaire,
				)
				id.POST(
					"/responses",
					middlewares.Authorize("write:responses"),
					h.CreateQuestionnaireResponse,
				)
				id.GET(
					"/responses",
					middlewares.Authorize("read:responses"),
					h.GetQuestionnaireResponses,
				)
			}
		}

		// TODO
		notifications := v1.Group("/notifications")
		{
			notifications.GET("", h.GetNotifications)
			notifications.GET("/:id", h.GetNotification)
			notifications.DELETE("/:id", h.DeleteNotification)
		}

		mycards := v1.Group("/mycards")
		mycards.Use(middlewares.Authorize("read:mycards"))
		{
			mycards.GET(
				"",
				middlewares.Authorize("read:mycards"),
				h.GetMyCards,
			)
			mycards.GET(
				"/:id",
				middlewares.Authorize("read:mycards"),
				h.GetMyCard,
			)
			mycards.DELETE(
				"/:id",
				middlewares.Authorize("write:mycards"),
				h.DeleteMyCard,
			)
		}

		prices := v1.Group("/prices")
		prices.Use(middlewares.Authorize("read:prices"))
		{
			prices.GET(
				"",
				middlewares.Authorize("read:prices"),
				h.GetPrices,
			)
			prices.GET(
				"/:id",
				middlewares.Authorize("read:prices"),
				h.GetPrice,
			)
			prices.DELETE(
				"/:id",
				middlewares.Authorize("write:prices"),
				h.DeletePrice,
			)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
