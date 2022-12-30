package users

import (
	"aircargo/internal/core/ports"
	"aircargo/pkg/logging"

	"github.com/gin-gonic/gin"
)

type RoutesHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Hey(ctx *gin.Context)

	Signin(ctx *gin.Context)
}

type Handler struct {
	logger       logging.Logger
	usersService ports.UsersService
}

func NewHandler(
	logger logging.Logger,
	usersService ports.UsersService,

) *Handler {
	return &Handler{
		logger:       logger,
		usersService: usersService,
	}
}
