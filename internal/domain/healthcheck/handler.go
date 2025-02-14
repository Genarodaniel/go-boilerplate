package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckInterface interface {
	HealthCheck(c *gin.Context)
}

type Healthcheck struct {
}

func NewHealthcheckHandler() HealthCheckInterface {
	return &Healthcheck{}
}

func (h *Healthcheck) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Health!"})
}
