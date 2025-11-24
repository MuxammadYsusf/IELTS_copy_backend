package controller

import grpc "github/http/copy/task3/grpc/client"

type Controller struct {
	GRPCClient grpc.ServiceManager
}

func NewController(grpcClient grpc.ServiceManager) *Controller {
	return &Controller{
		GRPCClient: grpcClient,
	}
}
