package adder

import "context"

// GRPCServer ...
type GRPCServer struct {
}

// Add interface of grpc
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
