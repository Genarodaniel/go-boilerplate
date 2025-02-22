package middleware

import (
	"context"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/server/routeutils"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace the request's context with the new one
		c.Request = c.Request.WithContext(ctx)

		// Create a channel to signal when the request is done
		done := make(chan struct{})

		// Run the request in a goroutine
		go func() {
			c.Next() // Process the request
			close(done)
		}()

		// Select between completion and timeout
		select {
		case <-ctx.Done(): // Timeout reached
			routeutils.HandleError(c, customerror.NewTimeoutError("request timeout"))
			return
		case <-done: // Request completed before timeout
			return
		}
	}
}
