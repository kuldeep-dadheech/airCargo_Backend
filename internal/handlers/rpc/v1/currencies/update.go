package currencies

import (
	pb "aircargo/proto/gen/app/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Update(ctx context.Context, req *pb.UpdateCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
