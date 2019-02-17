package server

import (
	"context"
	pb "supply/supply/grpc/proto"
	"supply/supply/ordering"
	"time"
)

type OrderingServer struct {
	svc ordering.OrderingService
}

// status.Errorf(codes.OK, "error: %s")

func (s *OrderingServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	err := s.svc.CreateOrder(in.Id, in.ProjectId)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return &pb.CreateOrderResponse{}, nil
}

func (s *OrderingServer) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.AddOrderItemResponse, error) {
	err := s.svc.AddOrderItem(in.Id, in.ProductId, in.Name, in.Uom)
	if err != nil {
		return &pb.AddOrderItemResponse{}, err
	}
	return &pb.AddOrderItemResponse{}, nil
}

func (s *OrderingServer) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	err := s.svc.RemoveOrderItem(in.Id, in.ProductId)
	if err != nil {
		return &pb.RemoveOrderItemResponse{}, err
	}
	return &pb.RemoveOrderItemResponse{}, nil
}

func (s *OrderingServer) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	err := s.svc.ModifyRequestedQuantity(in.Id, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ModifyRequestedQuantityResponse{}, err
	}
	return &pb.ModifyRequestedQuantityResponse{}, nil
}

func (s *OrderingServer) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	err := s.svc.SendOrder(in.Id)
	if err != nil {
		return &pb.SendOrderResponse{}, err
	}
	return &pb.SendOrderResponse{}, nil
}

func (s *OrderingServer) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	err := s.svc.ReceiveOrderItem(in.Id, in.ProductId, uint(in.Quantity))
	if err != nil {
		return &pb.ReceiveOrderItemResponse{}, err
	}
	return &pb.ReceiveOrderItemResponse{}, nil
}

func (s *OrderingServer) FindOrder(ctx context.Context, in *pb.FindOrderRequest) (*pb.FindOrderResponse, error) {
	order, err := s.svc.FindOrder(in.Id)
	if err != nil {
		return &pb.FindOrderResponse{}, err
	}

	var items []*pb.Item
	for _, i := range order.Items {
		item := &pb.Item{
			ProductId:         i.ProductID,
			Name:              i.Name,
			Uom:               i.UOM,
			QuantityRequested: uint32(i.QuantityRequested),
			QuantityReceived:  uint32(i.QuantityReceived),
			ItemStatus:        i.ItemStatus.String(),
		}
		items = append(items, item)
	}

	return &pb.FindOrderResponse{
		Date:   order.SentDate,
		Status: order.Status.String(),
		Items:  items,
	}, nil
}

func (s *OrderingServer) FindProjectOrderDates(ctx context.Context, in *pb.FindProjectOrderDatesRequest) (*pb.FindProjectOrderDatesResponse, error) {
	oo, err := s.svc.FindProjectOrderDates(in.ProjectId)
	if err != nil {
		return &pb.FindProjectOrderDatesResponse{}, err
	}

	var orders []*pb.Order
	for _, o := range oo {
		order := &pb.Order{
			Id:     o.ID,
			Date:   o.SentDate,
			Status: o.Status.String(),
		}
		orders = append(orders, order)
	}

	time.Sleep(2 * time.Second)

	return &pb.FindProjectOrderDatesResponse{Orders: orders}, nil
}
