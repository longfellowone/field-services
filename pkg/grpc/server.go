package server

import (
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/ordering"
	"supply/pkg/purchasing"
)

func New(os ordering.OrderingService, ps purchasing.PurchasingService) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterOrderingServer(s, &OrderingServer{
		svc: os,
	})
	return s
}
