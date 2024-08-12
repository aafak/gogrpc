package handlers

import (
	"context"
	"log"

	userpb "github.com/aafak/gogrpc/protos/user"
)

type UserServicer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServicer) CreateUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Printf("CreateUser function was invoked with request: %+v", req)
	return &userpb.UserResponse{Name: "Hello " + req.GetName(), Id: "123"}, nil
}
