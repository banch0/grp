package adder

import (
	"context"
	"grp/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
}

// Service provides password hashing capabilities.
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password, hash string) (bool, error)
}

// Add interface of grpc
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

// Check code ...
func (s *GRPCServer) Check(ctx context.Context, req *api.Code) (*api.Link, error) {
	return &api.Link{Id: req.GetId(), Code: req.GetLink()}, nil
}

// Send code ...
func (s *GRPCServer) Send(ctx context.Context, req *api.Link) (*api.Code, error) {
	return &api.Code{Id: req.GetId(), Link: req.GetCode()}, nil
}

// Hash ...
func (s *GRPCServer) Hash(ctx context.Context, req *api.HashRequest) (*api.HashResponse, error) {
	return &api.HashResponse{Hash: req.GetPassword()}, nil
}

// Validate ...
func (s *GRPCServer) Validate(ctx context.Context, req *api.ValidateRequest) (*api.ValidateResponse, error) {
	return &api.ValidateResponse{Valid: req.GetPassword() == req.GetPassword()}, nil
}
