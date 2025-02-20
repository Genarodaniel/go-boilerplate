package user

import (
	"go-boilerplate/internal/app/user"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, service *user.UserService) {
	handler := NewUserHandler(service)

	g.POST("/", handler.HandleRegister)
	g.GET("/:id", handler.HandleGet)
}
