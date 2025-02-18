package user

import (
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, service *UserService) {
	handler := NewUserHandler(service)

	g.POST("/", handler.HandlePostUser)
	g.GET("/:id", handler.HandleGetUser)
}
