package auth

import (
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/helper"
	"go-boilerplate/server/routeutils"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareInterface interface {
	AuthRequired(c *gin.Context)
}

type AuthMiddleware struct {
	jwtHelper helper.JWTHelper
}

func NewAuthMiddleware(jwtHelper helper.JWTHelper) AuthMiddlewareInterface {
	return &AuthMiddleware{
		jwtHelper: jwtHelper,
	}
}

func (a *AuthMiddleware) AuthRequired(c *gin.Context) {
	tokenFromHeader := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
	if tokenFromHeader == "" {
		routeutils.HandleError(c, customerror.NewUnauthorizedError("authorization token is required"))
		return
	}

	claims, err := a.jwtHelper.GetClaims(tokenFromHeader)
	if err != nil {
		routeutils.HandleError(c, customerror.NewUnauthorizedError(err.Error()))
		return
	}

	if claims == nil {
		routeutils.HandleError(c, customerror.NewUnauthorizedError("invalid authorization token"))
		return
	}
}
