package handlers

import (
	"context"
	"log"

	userpb "github.com/aafak/gogrpc/protos/user"
)

// UserServicer is a struct that implements the UserServiceServer interface
type UserServicer struct {
	userpb.UnimplementedUserServiceServer
}

// CreateUser is a function that implements the CreateUser method of the UserServiceServer interface
func (s *UserServicer) CreateUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Printf("CreateUser function was invoked with request: %+v", req)
	return &userpb.UserResponse{Name: "Hello " + req.GetName(), Id: "123"}, nil
}
