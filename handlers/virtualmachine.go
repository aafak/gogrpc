package handlers

import (
	"context"
	"log"

	vmpb "github.com/aafak/gogrpc/protos/virtualmachine"
)

type VirtualMachineServicer struct {
	vmpb.UnimplementedVirtualMachineServiceServer
}

func (s *VirtualMachineServicer) CreateVirtualMachine(ctx context.Context, req *vmpb.VirtualMachineRequest) (*vmpb.VirtualMachineResponse, error) {
	log.Printf("CreateVM function was invoked with request: %+v", req)
	return &vmpb.VirtualMachineResponse{Name: "Hello " + req.GetName(), Id: "123"}, nil
}
