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

	pb.RegisterSupplyServer(s, &Server{
		osvc: osvc,
		ssvc: ssvc,
	})

	return s
}
