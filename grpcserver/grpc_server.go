package main

import (
	"log"
	"net"

	"github.com/aafak/gogrpc/handlers"
	userpb "github.com/aafak/gogrpc/protos/user"
	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &handlers.UserServicer{})
	vmpb.RegisterVirtualMachineServiceServer(s, &handlers.VirtualMachineServicer{})

	log.Printf("Server  listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
