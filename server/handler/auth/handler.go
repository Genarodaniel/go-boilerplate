package auth

import (
	"go-boilerplate/internal/app/auth"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/server/routeutils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type AuthHandlerInterface interface {
	HandleAuth(ctx *gin.Context)
}

type AuthHandler struct {
	AuthService auth.AuthServiceInterface
}

func NewAuthHandler(authService auth.AuthServiceInterface) *AuthHandler {
	validate = validator.New()
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) HandleAuth(ctx *gin.Context) {
	request := model.OAuthRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		routeutils.HandleError(ctx, customerror.NewRequestError(model.ErrInvalidRequest.Error()))
		return
	}

	if err := validate.Struct(request); err != nil {
		routeutils.HandleError(ctx, customerror.NewValidationError(model.ErrRequiredCredentials.Error()))
		return
	}

	if request.GrantType != "client_credentials" {
		routeutils.HandleError(ctx, customerror.NewRequestError(model.ErrInvalidGrantType.Error()))
		return
	}

	resp, err := h.AuthService.Authenticate(ctx, request.ClientID, request.ClientSecret, request.GrantType)
	if err != nil {
		routeutils.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
