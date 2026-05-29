package server

import (
	"avadnet/internal/config"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const expectedHost = "localhost:8080"

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		cspPolicy := "default-src 'self'; connect-src *; font-src *; " +
			"script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';"
		c.Header("Content-Security-Policy", cspPolicy)
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		permPolicy := "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=()," +
			"magnetometer=(),gyroscope=(),fullscreen=(self),payment=()"
		c.Header("Permissions-Policy", permPolicy)
		c.Next()
	}
}

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	//Init router security headers
	router.Use(SecurityHeaders())
	if cfg.Env == "production" {
		router.Use(static.Serve("/", static.LocalFile("./public/", true)))
	}
	return router
}
