package main

import (
	"context"
	"log"
	"time"

	userpb "github.com/aafak/gogrpc/protos/user"
	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// PS C:\Users\aafakmoh\github-repos\personel\gogrpc> gofmt -s -w .\grpc_client.go

func main() {
	grpcClient, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not  connect: %v", err)
	}
	defer grpcClient.Close()
	userService := userpb.NewUserServiceClient(grpcClient)
	vmService := vmpb.NewVirtualMachineServiceClient(grpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Attach metadata to context
	md := metadata.Pairs(
		"authToken", "token000",
		"traceId", "tarceId123",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	userResponse, err := userService.CreateUser(ctx, &userpb.UserRequest{Name: "User1"})
	if err != nil {
		log.Panicf("could not create user: %v", err)
	}
	log.Printf("User Id: %s, Name: %s", userResponse.GetId(), userResponse.GetName())

	vmResponse, err := vmService.CreateVirtualMachine(ctx, &vmpb.VirtualMachineRequest{Name: "vm1"})
	if err != nil {
		log.Panicf("could not create VM: %v", err)
	}
	log.Printf("VM ID: %s, Name: %s", vmResponse.GetId(), vmResponse.GetName())
}
