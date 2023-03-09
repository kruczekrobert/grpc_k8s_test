package compute

import (
	context "context"
)

type Service struct {
	UnimplementedAddServiceServer
}

func (s *Service) Compute(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	return &AddResponse{
		Result: 1,
	}, nil
}
