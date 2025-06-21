package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCsrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfInHeader := c.Request.Header.Get("csrf-token")
		csrfInCookie, err := c.Cookie("csrf-token")

		if err != nil || csrfInHeader != csrfInCookie {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "token invalid",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
