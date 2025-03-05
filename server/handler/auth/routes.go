package auth

import (
	"go-boilerplate/internal/app/auth"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, authService *auth.AuthService) {
	handler := NewAuthHandler(authService)

	g.POST("/token", handler.HandleAuth)
}
