package main

import (
	"context"
	"log"
	"time"

	userpb "github.com/aafak/gogrpc/protos/user"
	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcClient.Close()
	userService := userpb.NewUserServiceClient(grpcClient)
	vmService := vmpb.NewVirtualMachineServiceClient(grpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := userService.CreateUser(ctx, &userpb.UserRequest{Name: "User1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("User Id: %s, Name: %s", r.GetId(), r.GetName())

	vm, err := vmService.CreateVirtualMachine(ctx, &vmpb.VirtualMachineRequest{Name: "vm1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("VM ID: %s, Name: %s", vm.GetId(), r.GetName())
}
