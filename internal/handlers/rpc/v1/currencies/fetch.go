package currencies

import (
	pb "aircargo/proto/gen/app/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Fetch(ctx context.Context, req *pb.FetchCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
