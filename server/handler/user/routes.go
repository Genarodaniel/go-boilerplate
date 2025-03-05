package user

import (
	"go-boilerplate/internal/app/user"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, service *user.UserService) {
	handler := NewUserHandler(service)

	g.POST("/register", handler.HandleRegister)
	g.GET("/", handler.HandleGet)
}
