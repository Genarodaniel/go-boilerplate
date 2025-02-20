package errhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *ValidationError:
		statusCode = http.StatusUnprocessableEntity // 422
	case *NotFoundError:
		statusCode = http.StatusNotFound // 404
	case *ApplicationError:
		statusCode = http.StatusInternalServerError // 500
	default:
		statusCode = http.StatusInternalServerError // Default to 500
	}

	c.JSON(statusCode, gin.H{"error": err.Error()})
}
