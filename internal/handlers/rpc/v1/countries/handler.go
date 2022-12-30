package countries

import (
	"sagebackend/internal/core/ports"
	"sagebackend/pkg/logging"

	pb "sagebackend/proto/gen/app/v1"
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
