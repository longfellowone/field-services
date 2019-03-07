package server

import (
	"field/supply/purchasing"
)

type PurchasingServer struct {
	svc purchasing.Service
}

//func (s *OrderingServer) UpdateItemPO(ctx context.Context, in *pb.UpdateItemPORequest) (*pb.UpdateItemPOResponse, error) {
//	err := s.svc.UpdateItemPO(in.OrderId, in.ProductId, in.Ponumber)
//	if err != nil {
//		return &pb.UpdateItemPOResponse{}, err
//	}
//	return &pb.UpdateItemPOResponse{}, nil
//}
