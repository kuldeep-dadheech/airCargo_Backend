package countries

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "sagebackend/proto/gen/app/v1"
)

func (h *Handler) Fetch(ctx context.Context, req *pb.FetchCountryRequest) (*pb.CountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
