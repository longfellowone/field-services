package server

import (
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/purchasing"
)

type PurchasingServer struct {
	svc purchasing.PurchasingService
}

func NewPurchasingServer(svr *grpc.Server, svc purchasing.PurchasingService) *PurchasingServer {
	pb.RegisterOrderingServer(svr, &OrderingServer{}) // Change

	return &PurchasingServer{
		svc: svc,
	}
}

//func (s *OrderingServer) UpdateItemPO(ctx context.Context, in *pb.UpdateItemPORequest) (*pb.UpdateItemPOResponse, error) {
//	err := s.svc.UpdateItemPO(in.OrderId, in.ProductId, in.Ponumber)
//	if err != nil {
//		return &pb.UpdateItemPOResponse{}, err
//	}
//	return &pb.UpdateItemPOResponse{}, nil
//}
