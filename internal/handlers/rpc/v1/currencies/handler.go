package currencies

import (
	"sagebackend/internal/core/ports"
	"sagebackend/pkg/logging"

	pb "sagebackend/proto/gen/app/v1"
)

type Handler struct {
	pb.UnimplementedCurrenciesServiceServer
	logger            logging.Logger
	currenciesService ports.CurrenciesService
}

func NewHandler(
	logger logging.Logger,
	currenciesService ports.CurrenciesService,
) *Handler {
	return &Handler{
		logger:            logger,
		currenciesService: currenciesService,
	}
}
