package server

import (
	"context"
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/ordering"
)

type Server struct {
	svc ordering.OrderingService
}

func NewOrderingServer(svr *grpc.Server, svc ordering.OrderingService) *Server {
	pb.RegisterOrderingServer(svr, &Server{})

	return &Server{
		svc: svc,
	}
}

func (s *Server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	err := s.svc.CreateOrder(in.OrderUuid, in.ProjectUuid)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return &pb.CreateOrderResponse{}, nil
}

func (s *Server) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.AddOrderItemResponse, error) {
	err := s.svc.AddOrderItem(in.OrderId, in.ProductId, in.Name, in.Uom)
	if err != nil {
		return &pb.AddOrderItemResponse{}, err
	}
	return &pb.AddOrderItemResponse{}, nil
}

func (s *Server) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	err := s.svc.RemoveOrderItem(in.OrderId, in.ProductId)
	if err != nil {
		return &pb.RemoveOrderItemResponse{}, err
	}
	return &pb.RemoveOrderItemResponse{}, nil
}

func (s *Server) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	err := s.svc.ModifyRequestedQuantity(in.OrderId, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ModifyRequestedQuantityResponse{}, err
	}
	return &pb.ModifyRequestedQuantityResponse{}, nil
}

func (s *Server) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	err := s.svc.SendOrder(in.OrderUuid)
	if err != nil {
		return &pb.SendOrderResponse{}, nil
	}
	return &pb.SendOrderResponse{}, nil
}

func (s *Server) UpdateItemPO(ctx context.Context, in *pb.UpdateItemPORequest) (*pb.UpdateItemPOResponse, error) {
	err := s.svc.UpdateItemPO(in.OrderId, in.ProductId, in.Ponumber)
	if err != nil {
		return &pb.UpdateItemPOResponse{}, nil
	}
	return &pb.UpdateItemPOResponse{}, nil
}

func (s *Server) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	err := s.svc.ReceiveOrderItem(in.OrderId, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ReceiveOrderItemResponse{}, nil
	}
	return &pb.ReceiveOrderItemResponse{}, nil
}
