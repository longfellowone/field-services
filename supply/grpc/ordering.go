package server

import (
	"context"
	pb "field/supply/grpc/proto"
)

// To return error status
// status.Errorf(codes.OK, "error: %s")

// Orders

func (s *Server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	_, err := s.osvc.CreateOrder(in.Id, in.ProjectId, in.Name, in.Foreman, in.Email)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return &pb.CreateOrderResponse{}, nil
}

func (s *Server) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.AddOrderItemResponse, error) {
	_, err := s.osvc.AddOrderItem(in.Id, in.ProductId, in.Name, in.Uom)
	if err != nil {
		return &pb.AddOrderItemResponse{}, err
	}
	return &pb.AddOrderItemResponse{}, nil
}

func (s *Server) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.RemoveOrderItemResponse, error) {
	_, err := s.osvc.RemoveOrderItem(in.Id, in.ProductId)
	if err != nil {
		return &pb.RemoveOrderItemResponse{}, err
	}
	return &pb.RemoveOrderItemResponse{}, nil
}

func (s *Server) ModifyRequestedQuantity(ctx context.Context, in *pb.ModifyRequestedQuantityRequest) (*pb.ModifyRequestedQuantityResponse, error) {
	_, err := s.osvc.ModifyRequestedQuantity(in.Id, in.ProductId, int(in.Quantity))
	if err != nil {
		return &pb.ModifyRequestedQuantityResponse{}, err
	}
	return &pb.ModifyRequestedQuantityResponse{}, nil
}

func (s *Server) SendOrder(ctx context.Context, in *pb.SendOrderRequest) (*pb.SendOrderResponse, error) {
	_, err := s.osvc.SendOrder(in.Id, in.Comments)
	if err != nil {
		return &pb.SendOrderResponse{}, err
	}
	return &pb.SendOrderResponse{}, nil
}

func (s *Server) ReceiveOrderItem(ctx context.Context, in *pb.ReceiveOrderItemRequest) (*pb.ReceiveOrderItemResponse, error) {
	_, err := s.osvc.ReceiveOrderItem(in.Id, in.ProductId, int(in.Quantity))
	if err != nil {
		return &pb.ReceiveOrderItemResponse{}, err
	}
	return &pb.ReceiveOrderItemResponse{}, nil
}

func (s *Server) FindOrder(ctx context.Context, in *pb.FindOrderRequest) (*pb.FindOrderResponse, error) {
	o, err := s.osvc.FindOrder(in.Id)
	if err != nil {
		return &pb.FindOrderResponse{}, err
	}
	//time.Sleep(2 * time.Second)

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
			QuantityRemaining: 0,
			ItemStatus:        i.ItemStatus.String(),
		}
		items = append(items, item)
	}

	order := &pb.Order{
		Id: o.ID,
		Project: &pb.Project{
			Id:   o.Project.ID,
			Name: o.Project.Name,
		},
		Items:    items,
		Date:     o.SentDate,
		Status:   o.Status.String(),
		Comments: o.Comments,
	}

	return &pb.FindOrderResponse{Order: order}, nil
}

func (s *Server) FindProjectOrderDates(ctx context.Context, in *pb.FindProjectOrderDatesRequest) (*pb.FindProjectOrderDatesResponse, error) {
	oo, err := s.osvc.FindProjectOrderDates(in.ProjectId)
	if err != nil {
		return &pb.FindProjectOrderDatesResponse{}, err
	}

	var orders []*pb.OrderSummary
	for _, o := range oo {
		order := &pb.OrderSummary{
			Id:     o.ID,
			Date:   o.SentDate,
			Status: o.Status.String(),
		}
		orders = append(orders, order)
	}

	return &pb.FindProjectOrderDatesResponse{Orders: orders}, nil
}

// Projects
func (s *Server) CloseProject(ctx context.Context, in *pb.CloseProjectRequest) (*pb.CloseProjectResponse, error) {
	return &pb.CloseProjectResponse{}, nil
}

func (s *Server) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	return &pb.CreateProjectResponse{}, nil
}

func (s *Server) FindProjects(ctx context.Context, in *pb.FindProjectsRequest) (*pb.FindProjectsResponse, error) {
	return &pb.FindProjectsResponse{}, nil
}
