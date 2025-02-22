package routeutils

import (
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/server/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *customerror.ValidationError:
		statusCode = http.StatusUnprocessableEntity
	case *customerror.NotFoundError:
		statusCode = http.StatusNotFound
	case *customerror.ApplicationError:
		statusCode = http.StatusInternalServerError
	case *customerror.UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *customerror.TimeoutError:
		statusCode = http.StatusRequestTimeout
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, response.DefaultErrorResponse{
		Error:   true,
		Message: err.Error(),
	})
}
