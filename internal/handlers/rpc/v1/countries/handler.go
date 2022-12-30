package countries

import (
	"aircargo/internal/core/ports"
	"aircargo/pkg/logging"

	pb "aircargo/proto/gen/app/v1"
)

type Handler struct {
	pb.UnimplementedCountriesServiceServer
	logger           logging.Logger
	countriesService ports.CountriesService
}

func NewHandler(
	logger logging.Logger,
	countriesService ports.CountriesService,
) *Handler {
	return &Handler{
		logger:           logger,
		countriesService: countriesService,
	}
}
