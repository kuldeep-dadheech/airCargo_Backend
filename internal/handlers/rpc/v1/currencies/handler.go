package currencies

import (
	"aircargo/internal/core/ports"
	"aircargo/pkg/logging"

	pb "aircargo/proto/gen/app/v1"
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
