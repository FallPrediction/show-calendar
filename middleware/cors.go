package middleware

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"content-type", "authorization", "csrf-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
