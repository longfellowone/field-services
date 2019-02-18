package search

import (
	"github.com/sahilm/fuzzy"
	"strings"
	"supply/supply"
)

type SearchService interface {
	ProductSearch(name string) []Result
}

type ProductRepository interface {
	FindAll() ([]supply.Product, error)
}

type Service struct {
	products []supply.Product
}

func NewSearchService(product ProductRepository) (*Service, error) {
	products, err := product.FindAll()
	if err != nil {
		return &Service{}, err
	}
	return &Service{products: products}, nil
}

func (s *Service) ProductSearch(name string) []Result {
	replacer := strings.NewReplacer("“", "\"", "”", "\"")
	output := replacer.Replace(name)

	fr := fuzzy.FindFrom(output, s)

	if fr.Len() > 10 {
		fr = fr[:10]
	}

	var results []Result
	for _, r := range fr {
		result := Result{
			Product:        s.products[r.Index],
			MatchedIndexes: r.MatchedIndexes,
		}
		results = append(results, result)
	}
	return results
}

// Result of ProductSearch
type Result struct {
	supply.Product
	MatchedIndexes []int
}

// Required methods for fuzzy search
func (s *Service) String(i int) string {
	return s.products[i].Name
}

func (s *Service) Len() int {
	return len(s.products)
}
