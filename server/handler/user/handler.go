package user

import (
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/app/user"
	"go-boilerplate/pkg/validation"
	"go-boilerplate/server/response"
	"go-boilerplate/server/routeutils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type UserHandlerInterface interface {
	HandleRegister(ctx *gin.Context)
	HandleGet(ctx *gin.Context)
}

type UserHandler struct {
	UserService user.UserServiceInterface
}

func NewUserHandler(UserService user.UserServiceInterface) *UserHandler {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return &UserHandler{
		UserService: UserService,
	}
}

func (h *UserHandler) HandleRegister(ctx *gin.Context) {
	request := &model.PostUserRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.UserService.Register(ctx, user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		routeutils.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *UserHandler) HandleGet(ctx *gin.Context) {
	userID := ctx.Param("id")
	if valid := validation.ValidateUUID(userID); !valid {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DefaultErrorResponse{
			Error:   true,
			Message: model.ErrInvalidUUID.Error(),
		})
		return
	}

	resp, err := h.UserService.Get(ctx, userID)
	if err != nil {
		routeutils.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
