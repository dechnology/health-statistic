package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true, // this allows all origins
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Accept",
			"User-Agent",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-length"},
	}

	return cors.New(c)
}
