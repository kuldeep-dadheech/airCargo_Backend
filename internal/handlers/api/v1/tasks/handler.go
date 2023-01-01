package tasks

import (
	"aircargo/internal/core/ports"
	"aircargo/pkg/logging"

	"github.com/gin-gonic/gin"
)


type RoutesHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type Handler struct {
	logger    logging.Logger
	taskService ports.TasksService
}

func NewHandler(
	logger logging.Logger,
	taskService ports.TasksService,
) *Handler {
	return &Handler{
		logger:    logger,
		taskService: taskService,
	}
}