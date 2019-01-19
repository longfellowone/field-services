package server

import (
	"context"
	pb "supply/api/grpc/proto"
	"supply/api/search"
)

type SearchServer struct {
	svc search.SearchService
}

func (s *SearchServer) ProductSearch(ctx context.Context, in *pb.ProductSearchRequest) *pb.ProductSearchResponse {
	products := s.svc.ProductSearch(in.Name)

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

	return &pb.ProductSearchResponse{Results: results}
}
