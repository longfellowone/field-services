package server

import (
	pb "field/supply/grpc/proto"
	"field/supply/ordering"
	"field/supply/purchasing"
	"field/supply/search"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type Server struct {
	osvc ordering.Service
	ssvc search.Service
}

func New(osvc ordering.Service, psvc purchasing.Service, ssvc search.Service) *grpc.Server {
	s := grpc.NewServer()

	_ = osvc.CreateOrder("oid1", "pid1")
	_ = osvc.CreateOrder("oid2", "pid1")
	_ = osvc.CreateOrder("oid3", "pid1")
	_ = osvc.CreateOrder("oid4", "pid2")
	_ = osvc.CreateOrder("oid5", "pid2")

	orders := osvc.FindProjectOrderDates("pid1")

	for _, o := range orders {
		fmt.Println(o.ID, time.Unix(o.SentDate, 0))
	}

	_ = osvc.AddOrderItem("oid1", "pid1", "product1", "ea")

	pb.RegisterSupplyServer(s, &Server{
		osvc: osvc,
		ssvc: ssvc,
	})

	return s
}
