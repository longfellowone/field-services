package search

import (
	"github.com/sahilm/fuzzy"
	"supply/pkg"
)

type SearchService interface {
	ProductSearch(query string) []Result
}

type ProductRepository interface {
	FindAll() ([]supply.Product, error)
}

type Service struct {
	products products
}

func NewSearchService(product ProductRepository) (*Service, error) {
	products, _ := product.FindAll()

	return &Service{products: products}, nil
}

func (s *Service) ProductSearch(query string) []Result {
	fr := fuzzy.FindFrom(query, s.products)

	if fr.Len() > 10 {
		fr = fr[:10]
	}

	var results []Result
	for i, r := range fr {
		result := Result{
			Product:        s.products[i],
			MatchedIndexes: r.MatchedIndexes,
		}
		results = append(results, result)
	}
	return results
}

type Result struct {
	Product        supply.Product
	MatchedIndexes []int
}

type products []supply.Product

func (p products) String(i int) string {
	return p[i].Name
}

func (p products) Len() int {
	return len(p)
}
