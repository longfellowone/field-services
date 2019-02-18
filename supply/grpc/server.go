package server

import (
	"fmt"
	"google.golang.org/grpc"
	pb "supply/supply/grpc/proto"
	"supply/supply/ordering"
	"supply/supply/purchasing"
	"supply/supply/search"
	"time"
)

type SupplyServer struct {
	osvc ordering.OrderingService
	ssvc search.SearchService
}

func New(osvc ordering.OrderingService, psvc purchasing.PurchasingService, ssvc search.SearchService) *grpc.Server {
	s := grpc.NewServer()

	osvc.CreateOrder("oid1", "pid1")
	osvc.CreateOrder("oid2", "pid1")
	osvc.CreateOrder("oid3", "pid1")
	osvc.CreateOrder("oid4", "pid2")
	osvc.CreateOrder("oid5", "pid2")

	orders, _ := osvc.FindProjectOrderDates("pid1")

	for _, o := range orders {
		fmt.Println(o.ID, time.Unix(o.SentDate, 0))
	}

	osvc.AddOrderItem("oid1", "pid1", "product1", "ea")

	pb.RegisterSupplyServer(s, &SupplyServer{
		osvc: osvc,
		ssvc: ssvc,
	})

	return s
}
