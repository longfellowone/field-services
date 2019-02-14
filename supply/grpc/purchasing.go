package server

import (
	"supply/supply/purchasing"
)

type PurchasingServer struct {
	svc purchasing.PurchasingService
}

//func (s *OrderingServer) UpdateItemPO(ctx context.Context, in *pb.UpdateItemPORequest) (*pb.UpdateItemPOResponse, error) {
//	err := s.svc.UpdateItemPO(in.OrderId, in.ProductId, in.Ponumber)
//	if err != nil {
//		return &pb.UpdateItemPOResponse{}, err
//	}
//	return &pb.UpdateItemPOResponse{}, nil
//}
