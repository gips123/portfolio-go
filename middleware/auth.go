package middleware

import (
	"best-portfolio-go/config"
	"best-portfolio-go/utils"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// RequireAuth enforces JWT auth for protected routes.
func RequireAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "missing or invalid Authorization header",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(token, cfg.Auth.JWTSecret)
		if err != nil {
			status := http.StatusUnauthorized
			if errors.Is(err, utils.ErrTokenExpired) {
				status = http.StatusUnauthorized
			}
			c.AbortWithStatusJSON(status, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Store claims for downstream handlers (if needed)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("token_exp", time.Unix(claims.Exp, 0))

		c.Next()
	}
}
