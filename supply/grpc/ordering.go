package server

import (
	"context"
	"field/supply"
	pb "field/supply/grpc/proto"
	"time"
)

// To return error status
// status.Errorf(codes.OK, "error: %s")

// Orders

func (s *Server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	o, err := s.osvc.CreateOrder(in.ProjectId, in.Name, in.Foreman, in.Email)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}

	os := &pb.OrderSummary{
		Id:     o.ID,
		Date:   int32(o.SentDate),
		Status: o.Status.String(),
	}
	return &pb.CreateOrderResponse{Order: os}, nil
}

func (s *Server) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.AddOrderItemResponse, error) {
	o, err := s.osvc.AddOrderItem(in.Id, in.ProductId, in.Name, in.Uom)
	if err != nil {
		return &pb.AddOrderItemResponse{}, err
	}
	return &pb.AddOrderItemResponse{Order: orderToProto(o)}, nil
}

func (s *Server) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	o, err := s.osvc.RemoveOrderItem(in.Id, in.ProductId)
	if err != nil {
		return &pb.RemoveOrderItemResponse{}, err
	}
	return &pb.RemoveOrderItemResponse{Order: orderToProto(o)}, nil
}

func (s *Server) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	o, err := s.osvc.ModifyRequestedQuantity(in.Id, in.ProductId, int(in.Quantity))
	if err != nil {
		return &pb.ModifyRequestedQuantityResponse{}, err
	}
	return &pb.ModifyRequestedQuantityResponse{Order: orderToProto(o)}, nil
}

func (s *Server) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	o, err := s.osvc.SendOrder(in.Id, in.Comments)
	if err != nil {
		return &pb.SendOrderResponse{}, err
	}
	return &pb.SendOrderResponse{Order: orderToProto(o)}, nil
}

func (s *Server) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	o, err := s.osvc.ReceiveOrderItem(in.Id, in.ProductId, int(in.Quantity))
	if err != nil {
		return &pb.ReceiveOrderItemResponse{}, err
	}
	return &pb.ReceiveOrderItemResponse{Order: orderToProto(o)}, nil
}

func (s *Server) FindOrder(ctx context.Context, in *pb.FindOrderRequest) (*pb.FindOrderResponse, error) {
	o, err := s.osvc.FindOrder(in.Id)
	if err != nil {
		return &pb.FindOrderResponse{}, err
	}

	//TEMPORARY
	time.Sleep(1 * time.Second)

	return &pb.FindOrderResponse{Order: orderToProto(o)}, nil
}

func (s *Server) FindProjectOrderDates(ctx context.Context, in *pb.FindProjectOrderDatesRequest) (*pb.FindProjectOrderDatesResponse, error) {
	oo, err := s.osvc.FindProjectOrderDates(in.ProjectId)
	if err != nil {
		return &pb.FindProjectOrderDatesResponse{}, err
	}

	// TEMPORARY
	time.Sleep(1 * time.Second)

	var orders []*pb.OrderSummary
	for _, o := range oo {
		order := &pb.OrderSummary{
			Id:          o.ID,
			Date:        int32(o.SentDate),
			ProjectName: o.ProjectName,
			Status:      o.Status.String(),
		}
		orders = append(orders, order)
	}

	return &pb.FindProjectOrderDatesResponse{Orders: orders}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, in *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	err := s.osvc.DeleteOrder(in.OrderId)
	if err != nil {
		return &pb.DeleteOrderResponse{}, err
	}
	return &pb.DeleteOrderResponse{}, nil
}

func orderToProto(o *supply.Order) *pb.Order {
	var items []*pb.Item
	for _, i := range o.Items {
		item := &pb.Item{
			Product: &pb.Product{
				Id:   i.ID,
				Name: i.Name,
				Uom:  i.UOM,
			},
			QuantityRequested: uint32(i.QuantityRequested),
			QuantityReceived:  uint32(i.QuantityReceived),
			QuantityRemaining: uint32(i.QuantityRemaining),
			ItemStatus:        i.ItemStatus.String(),
			Deleted:           i.Removed,
		}
		items = append(items, item)
	}

	return &pb.Order{
		Id: o.ID,
		Project: &pb.Project{
			Id:   o.Project.ID,
			Name: o.Project.Name,
		},
		Items:    items,
		Date:     int32(o.SentDate),
		Status:   o.Status.String(),
		Comments: o.Comments,
	}
}

// Projects
func (s *Server) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	return &pb.CreateProjectResponse{}, nil
}

func (s *Server) CloseProject(ctx context.Context, in *pb.CloseProjectRequest) (*pb.CloseProjectResponse, error) {
	return &pb.CloseProjectResponse{}, nil
}

func (s *Server) FindProjects(ctx context.Context, in *pb.FindProjectsRequest) (*pb.FindProjectsResponse, error) {
	// TEMPORARY
	//time.Sleep(500 * time.Millisecond)

	return &pb.FindProjectsResponse{
		Projects: []*pb.Project{{
			Id:   "1",
			Name: "Test1",
		}, {
			Id:   "2",
			Name: "Test2",
		}},
	}, nil
}
