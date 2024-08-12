package handlers

import (
	"context"
	"log"

	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"
)

// VirtualMachineServicer is a struct that implements the VirtualMachineServiceServer interface
type VirtualMachineServicer struct {
	vmpb.UnimplementedVirtualMachineServiceServer
}

// CreateVirtualMachine is a function that implements the CreateVirtualMachine method of the VirtualMachineServiceServer interface
func (s *VirtualMachineServicer) CreateVirtualMachine(ctx context.Context, req *vmpb.VirtualMachineRequest) (*vmpb.VirtualMachineResponse, error) {
	log.Printf("CreateVM function was invoked with request: %+v", req)
	return &vmpb.VirtualMachineResponse{Name: "Hello " + req.GetName(), Id: "123"}, nil
}
