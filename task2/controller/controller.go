package controller

import grpc "github/http/copy/task2/grpc/client"

type Handler struct {
	GRPCClient grpc.ServiceManager
}

func Controller(grpcClinet grpc.ServiceManager) *Handler {

	return &Handler{
		GRPCClient: grpcClinet,
	}
}
