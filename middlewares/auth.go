package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user")
		if sessionID == nil {
			c.HTML(http.StatusOK, "index", gin.H{
				"title": "consumer reviews",
				"user":  sessionID,
			})
			c.Abort()
		}
	}
}
