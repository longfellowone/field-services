package server

import (
	"fmt"
	"google.golang.org/grpc"
	pb "supply/api/grpc/proto"
	"supply/api/ordering"
	"supply/api/purchasing"
	"supply/api/search"
	"time"
)

func New(osvc ordering.OrderingService, psvc purchasing.PurchasingService, ssvc search.SearchService) *grpc.Server {
	s := grpc.NewServer()

	osvc.CreateOrder("oid1", "pid1")
	osvc.CreateOrder("oid2", "pid1")
	osvc.CreateOrder("oid3", "pid1")
	osvc.CreateOrder("oid4", "pid2")
	osvc.CreateOrder("oid5", "pid2")

	orders, _ := osvc.FindProjectOrderDates("pid1")

	for _, o := range orders {
		fmt.Println(o.OrderID, time.Unix(o.SentDate, 0))
	}

	osvc.AddOrderItem("oid1", "pid1", "product1", "ea")

	pb.RegisterOrderingServer(s, &OrderingServer{svc: osvc})
	//pb.RegisterOrderingServer(s, &PurchasingServer{svc: psvc})
	pb.RegisterSearchServer(s, &SearchServer{svc: ssvc})

	return s
}
