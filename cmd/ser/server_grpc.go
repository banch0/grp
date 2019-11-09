package main

import (
	"grp/pkg/api"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, r *api.HashRequest) (*api.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*api.HashResponse), nil
}

func (s *grpcServer) Validte(ctx context.Context, r *api.ValidateRequest) (*api.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*api.ValidateResponse), nil
}

// EncodeGRPCHashRequest ...
func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &api.HashRequest{Password: req.Password}, nil
}

//DecodeGRPCHashRequest ...
func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*api.HashRequest)
	return hashRequest{Password: req.Password}, nil
}

// EncodeGRPCHashResponse ...
func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &api.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

// DecodeGRPCHashResponse ...
func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*api.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}

// EncodeGRPCValidateRequest ...
func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &api.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

// DecodeGRPCValidateRequest ...
func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*api.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}

// EncodeGRPCValidateResponse ...
func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &api.ValidateResponse{Valid: res.Valid}, nil
}
