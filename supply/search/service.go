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
	FindAll() ([]supply.Product, error)
}

type Svc struct {
	products []supply.Product
}

func NewSearchService(product productRepository) *Svc {
	products, err := product.FindAll()
	if err != nil {
		panic(err)
	}
	return &Svc{products: products}
}

func (s *Svc) ProductSearch(name string) []Result {
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
func (s *Svc) String(i int) string {
	return s.products[i].Name
}

func (s *Svc) Len() int {
	return len(s.products)
}
