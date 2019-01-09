package grpc

import (
	"context"
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
)

type OrderingService interface {
	CreateOrder(orderid, projectid string) error
	AddOrderItem(orderid, productid, name, uom string) error
	RemoveOrderItem(orderid, productid string) error
	ModifyRequestedQuantity(orderid, productid string, quantity uint) error
	SendOrder(orderid string) error
	UpdateItemPO(orderid, productid, ponumber string) error
	ReceiveOrderItem(orderid, productid string, quantity uint) error
}

type Server struct {
	svc OrderingService
}

func NewOrderingServer(svr *grpc.Server, svc OrderingService) *Server {
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
	return &pb.AddOrderItemResponse{}, nil
}

func (s *Server) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	return &pb.RemoveOrderItemResponse{}, nil
}

func (s *Server) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	return &pb.ModifyRequestedQuantityResponse{}, nil
}

func (s *Server) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	return &pb.SendOrderResponse{}, nil
}

func (s *Server) UpdateItemPO(ctx context.Context, in *pb.UpdateItemPORequest) (*pb.UpdateItemPOResponse, error) {
	return &pb.UpdateItemPOResponse{}, nil
}

func (s *Server) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	return &pb.ReceiveOrderItemResponse{}, nil
}
