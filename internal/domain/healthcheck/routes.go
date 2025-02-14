package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	healthcheckHandler := NewHealthcheckHandler()
	g.GET("/", healthcheckHandler.HealthCheck)
}
