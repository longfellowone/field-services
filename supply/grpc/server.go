//go:generate protoc -I proto/ proto/supply.proto --go_out=plugins=grpc:proto
package server

import (
	pb "field/supply/grpc/proto"
	"field/supply/ordering"
	"field/supply/search"
	"google.golang.org/grpc"
)

type Server struct {
	osvc ordering.Service
	ssvc search.Service
}

func New(osvc ordering.Service, ssvc search.Service) *grpc.Server {
	s := grpc.NewServer()
	server := &Server{
		osvc: osvc,
		ssvc: ssvc,
	}
	pb.RegisterSupplyServer(s, server)
	return s
}

//request := &pb.FindOrderRequest{Id: "7e55aa12-2e6a-4f21-b01a-09503c755180"}
//response, err := server.FindOrder(context.Background(), request)
//if err != nil {
//	log.Fatal(err)
//}
//fmt.Println(response.Order.Id)
