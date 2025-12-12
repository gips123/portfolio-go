package handlers

import (
	"best-portfolio-go/config"
	"best-portfolio-go/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler returns a gin.HandlerFunc that issues JWT tokens
func LoginHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
			return
		}

		if req.Username != cfg.Auth.AdminUsername || req.Password != cfg.Auth.AdminPassword {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "invalid username or password",
			})
			return
		}

		ttl := time.Duration(cfg.Auth.TokenTTLMinutes) * time.Minute
		token, err := utils.GenerateToken(cfg.Auth.JWTSecret, req.Username, "admin", ttl)
		if err != nil {
			utils.InternalServerErrorResponse(c, "Failed to generate token: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"token":      token,
				"expiresIn":  int(ttl.Seconds()),
				"tokenType":  "Bearer",
				"issuedAt":   time.Now().UTC(),
				"expiresAt":  time.Now().UTC().Add(ttl),
				"user":       req.Username,
				"userRole":   "admin",
				"tokenScope": "crud",
			},
		})
	}
}
