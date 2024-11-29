package middleware

import (
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowOrigins: []string{"http://192.168.1.126:3000", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		// ExposeHeaders:    []string{"Content-Length", "Authorization", "token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://192.168.1.126:3000" || origin == "http://localhost:3000" {
				return true
			}
			return false
			// return true
		},
		MaxAge: 12 * time.Hour,
	}
	return cors.New(cfg)

}
