package server

import (
	"context"
	pb "supply/supply/grpc/proto"
)

//var (
//	ErrNoResults = errors.New("no results")
//)

func (s *SupplyServer) ProductSearch(ctx context.Context, in *pb.ProductSearchRequest) (*pb.ProductSearchResponse, error) {
	if in.Name == "" {
		return &pb.ProductSearchResponse{}, nil
	}

	products := s.ssvc.ProductSearch(in.Name)
	if len(products) == 0 {
		return &pb.ProductSearchResponse{}, nil
	}

	var results []*pb.Result
	for _, p := range products {

		var indexes []*pb.Index
		for _, i := range p.MatchedIndexes {
			index := &pb.Index{
				Index: int32(i),
			}
			indexes = append(indexes, index)
		}

		result := &pb.Result{
			ProductUuid: p.ProductID,
			Category:    p.Category,
			Name:        p.Name,
			Uom:         p.UOM,
			Indexes:     indexes,
		}
		results = append(results, result)
	}

	return &pb.ProductSearchResponse{Results: results}, nil
}
