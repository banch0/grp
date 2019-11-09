package sever

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-kit/kit/endpoint"
)

// Service ...
type Service struct{}

// NewService ...
func NewService() Service {
	return vaultService{}
}

// vaultService ...
type vaultService struct{}

// Hash ...
func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	log.Println("hash")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Validate ...
func (vaultService) Validate(ctx context.Context, password, hash string) (bool, error) {
	log.Println("validate")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

type hashRequest struct {
	Password string `json:"password"`
}

type hashResponse struct {
	Hash  string `json:"hash"`
	Error string `json:"err, omitempty"`
}

type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err, omitempty"`
}

func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("decode hash request")
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("decode validate request")
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func MakeHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{})
	(interface{}, error) {
		req := request.(hashRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return hashResponse{v, err.Error()}, err
		}
		return hashResponse{v, ""}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{})
	(interface{}, error) {
		req := request.(validateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse{false, err.Error()},nil
		}
		return validateResponse{v, ""}, nil
	}
}

type Endpoints struct {
	HashEndpoint endoint.Endpoint
	ValidataHashEndpoint endoint.Endpoint
}

func (e Endpoints) Hash(ctx context.Context, password string) (string, error) { 
	req := hashRequest{Password: password}
	resp, err := e.HashEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	hashResp := resp.(hashResponse)
	if hashResp.Err != "" {
		return "", errors.New(hashResp.Err)
	}
	return hashResp.Hash, nil
}

func (e Endpoints) Validate(ctx context.Context, password, 
	hash string)(bool, error) {
		req := validateRequest{Password: password, Hash: hash}
		resp, err := e.ValidataHashEndpoint(ctx, req)
		if err != nil {
			return false, err
		}
		validataResp := resp.(validateResponse)
		if validateResp.Err != "" {
			return false, errors.New(validataResp.err)
		}
		return validataResp.Valid, nil
	}