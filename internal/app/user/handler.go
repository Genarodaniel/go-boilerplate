package user

import (
	"go-boilerplate/internal/server/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	HandlePostUser(ctx *gin.Context)
}

type UserHandler struct {
	UserService UserServiceInterface
}

func NewUserHandler(UserService UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: UserService,
	}
}

func (h *UserHandler) HandlePostUser(ctx *gin.Context) {
	request := &PostUserRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	if err := request.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.UserService.PostUser(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
