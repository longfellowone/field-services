package server

import (
	"google.golang.org/grpc"
	pb "supply/pkg/grpc/proto"
	"supply/pkg/ordering"
	"supply/pkg/purchasing"
	"supply/pkg/search"
)

func New(osvc ordering.OrderingService, psvc purchasing.PurchasingService, ssvc search.SearchService) *grpc.Server {
	s := grpc.NewServer()

	pb.RegisterOrderingServer(s, &OrderingServer{svc: osvc})
	//pb.RegisterOrderingServer(s, &PurchasingServer{svc: psvc})
	//pb.RegisterOrderingServer(s, &SearchServer{svc: psvc})

	return s
}
