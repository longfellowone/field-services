package search

import (
	"field/supply"
	"github.com/sahilm/fuzzy"
	"strings"
)

type Service interface {
	ProductSearch(name string) []Result
}

type productRepository interface {
	FindAll() ([]*supply.Product, error)
}

type service struct {
	products []*supply.Product
}

func NewSearchService(product productRepository) *service {
	products, err := product.FindAll()
	if err != nil {
		panic(err)
	}
	return &service{products: products}
}

func (s *service) ProductSearch(name string) []Result {
	replacer := strings.NewReplacer("“", "\"", "”", "\"")
	output := replacer.Replace(name)

	fr := fuzzy.FindFrom(output, s)

	if fr.Len() > 10 {
		fr = fr[:10]
	}

	var results []Result
	for _, r := range fr {
		result := Result{
			ID:             s.products[r.Index].ID,
			Name:           s.products[r.Index].Name,
			UOM:            s.products[r.Index].UOM,
			MatchedIndexes: r.MatchedIndexes,
		}
		results = append(results, result)
	}
	return results
}

// Result of ProductSearch
type Result struct {
	ID             string
	Name           string
	UOM            string
	MatchedIndexes []int
}

// Required methods for fuzzy search
func (s *service) String(i int) string {
	return s.products[i].Name
}

func (s *service) Len() int {
	return len(s.products)
}
