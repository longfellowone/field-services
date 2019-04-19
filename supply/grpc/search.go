package server

import (
	"context"
	pb "field/supply/grpc/proto"
)

//var (
//	ErrNoResults = errors.New("no results")
//)

func (s *Server) ProductSearch(ctx context.Context, in *pb.ProductSearchRequest) (*pb.ProductSearchResponse, error) {
	if in.Name == "" {
		return &pb.ProductSearchResponse{}, nil
	}
	//time.Sleep(2000 * time.Millisecond)

	products, err := s.ssvc.ProductSearch(in.Name)
	if err != nil {
		return &pb.ProductSearchResponse{}, nil // TODO: handle error
	}
	if len(products) == 0 {
		return &pb.ProductSearchResponse{}, nil
	}

	var results []*pb.Result
	for _, p := range products {

		var indexes []int32
		for _, i := range p.MatchedIndexes {
			indexes = append(indexes, int32(i))
		}

		result := &pb.Result{
			Product: &pb.Product{
				Id:   p.ID,
				Name: p.Name,
				Uom:  p.UOM,
			},
			Indexes: indexes,
		}
		results = append(results, result)
	}

	return &pb.ProductSearchResponse{Results: results}, nil
}
