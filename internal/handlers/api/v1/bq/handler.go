package bq

import (
	"sagebackend/pkg/logging"
 "sagebackend/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type RoutesHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type Handler struct {
	logger logging.Logger
	bqService ports.BqService
}

func NewHandler(
	logger logging.Logger,
	bqService ports.BqService,
) *Handler {
	return &Handler{
		logger: logger,
		bqService: bqService,
	}
}