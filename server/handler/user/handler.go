package user

import (
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/app/user"
	"go-boilerplate/internal/infra/errhandler"
	"go-boilerplate/pkg/validation"
	"go-boilerplate/server/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type UserHandlerInterface interface {
	HandlePostUser(ctx *gin.Context)
	HandleGetUser(ctx *gin.Context)
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

func (h *UserHandler) HandlePostUser(ctx *gin.Context) {
	request := &model.PostUserRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	if validate.Struct(request) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultErrorResponse{
			Error:   true,
			Message: "name and email are required",
		})
		return
	}

	resp, err := h.UserService.PostUser(ctx, user.User{
		Name:  request.Name,
		Email: request.Email,
	})

	if err != nil {
		errhandler.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *UserHandler) HandleGetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if valid := validation.IsUUID(userID); !valid {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DefaultErrorResponse{
			Error:   true,
			Message: model.ErrInvalidUUID.Error(),
		})
		return
	}

	resp, err := h.UserService.GetUser(ctx, userID)
	if err != nil {
		errhandler.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
