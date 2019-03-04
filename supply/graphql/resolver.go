//go:generate go run scripts/gqlgen.go -v
package graphql

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"net/http"

	//. "field/supply/graphql/models"
	"field/supply/search"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	//osvc ordering.OrderingService
	ssvc search.SearchService
}

func New(ssvc search.SearchService) http.HandlerFunc {
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{
		//osvc: osvc,
		ssvc: ssvc,
	}}))
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Products(ctx context.Context, name string) ([]search.Result, error) {
	return r.ssvc.ProductSearch(name), nil
}
