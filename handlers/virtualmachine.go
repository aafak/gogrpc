package handlers

import (
	"context"
	"fmt"
	"log"

	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"
	"google.golang.org/grpc/metadata"
)

// VirtualMachineServicer is a struct that implements the VirtualMachineServiceServer interface
type VirtualMachineServicer struct {
	vmpb.UnimplementedVirtualMachineServiceServer
}

// CreateVirtualMachine is a function that implements the CreateVirtualMachine method of the VirtualMachineServiceServer interface
func (s *VirtualMachineServicer) CreateVirtualMachine(ctx context.Context, req *vmpb.VirtualMachineRequest) (*vmpb.VirtualMachineResponse, error) {
	log.Printf("CreateVM function was invoked with request: %+v", req)

	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("Metadata: %v, ok: %v\n", md, ok)
	if !ok {
		log.Println("No metadata found")
	}

	// Access specific metadata values
	authTokenValues := md.Get("authToken") //  [token1]
	if len(authTokenValues) != 0 {
		fmt.Printf("authToken: %v\n", authTokenValues[0])
	}
	traceId := md.Get("traceId")
	fmt.Printf("traceId: %v, type: %T\n", traceId, traceId) // []string

	spanId := md.Get("spanId")                           // invalid key
	fmt.Printf("spanId: %v, type: %T\n", spanId, spanId) //  [], type: []string

	return &vmpb.VirtualMachineResponse{Name: "Hello " + req.GetName(), Id: "123"}, nil
}
