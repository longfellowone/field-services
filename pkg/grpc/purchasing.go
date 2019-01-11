package server

import (
	"supply/pkg/purchasing"
)

type PurchasingServer struct {
	svc purchasing.PurchasingService
}

func NewPurchasingServer(svc purchasing.PurchasingService) *PurchasingServer {
	//pb.RegisterPurchasingServer(svr, &PurchasingServer{})

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
