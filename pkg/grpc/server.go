package server

import (
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/ordering"
)

func New(os ordering.OrderingService) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterOrderingServer(s, &OrderingServer{
		svc: os,
	})
	return s
}
