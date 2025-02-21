package routeutils

import (
	"go-boilerplate/internal/infra/customerror"
	"go-boilerplate/server/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *customerror.ValidationError:
		statusCode = http.StatusUnprocessableEntity // 422
	case *customerror.NotFoundError:
		statusCode = http.StatusNotFound // 404
	case *customerror.ApplicationError:
		statusCode = http.StatusInternalServerError // 500
	default:
		statusCode = http.StatusInternalServerError // Default to 500
	}

	c.JSON(statusCode, response.DefaultErrorResponse{
		Error:   true,
		Message: err.Error(),
	})
}
