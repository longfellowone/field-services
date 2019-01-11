package server

import (
	"context"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/ordering"
)

type OrderingServer struct {
	svc ordering.OrderingService
}

func NewOrderingServer(svc ordering.OrderingService) *OrderingServer {
	return &OrderingServer{
		svc: svc,
	}
}

func (s *OrderingServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	err := s.svc.CreateOrder(in.OrderUuid, in.ProjectUuid)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return &pb.CreateOrderResponse{}, nil
}

func (s *OrderingServer) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.AddOrderItemResponse, error) {
	err := s.svc.AddOrderItem(in.OrderId, in.ProductId, in.Name, in.Uom)
	if err != nil {
		return &pb.AddOrderItemResponse{}, err
	}
	return &pb.AddOrderItemResponse{}, nil
}

func (s *OrderingServer) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	err := s.svc.RemoveOrderItem(in.OrderId, in.ProductId)
	if err != nil {
		return &pb.RemoveOrderItemResponse{}, err
	}
	return &pb.RemoveOrderItemResponse{}, nil
}

func (s *OrderingServer) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	err := s.svc.ModifyRequestedQuantity(in.OrderId, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ModifyRequestedQuantityResponse{}, err
	}
	return &pb.ModifyRequestedQuantityResponse{}, nil
}

func (s *OrderingServer) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	err := s.svc.SendOrder(in.OrderUuid)
	if err != nil {
		return &pb.SendOrderResponse{}, err
	}
	return &pb.SendOrderResponse{}, nil
}

func (s *OrderingServer) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	err := s.svc.ReceiveOrderItem(in.OrderId, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ReceiveOrderItemResponse{}, err
	}
	return &pb.ReceiveOrderItemResponse{}, nil
}
