//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/longfellowone/field-services/supply/ordering"
	"github.com/longfellowone/field-services/supply/search"
	"net/http"

	"github.com/longfellowone/field-services/supply"
	"github.com/longfellowone/field-services/supply/graphql/models"
)

type Resolver struct {
	osvc ordering.Service
	ssvc search.Service
}

func Initialize(ssvc search.Service, osvc ordering.Service) http.HandlerFunc {
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{
		osvc: osvc,
		ssvc: ssvc,
	}}))
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrder) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) SendOrder(ctx context.Context, input models.SendOrder) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddOrderItem(ctx context.Context, input models.AddOrderItem) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) RemoveOrderItem(ctx context.Context, input models.RemoveOrderItem) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) ReceiveOrderItem(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) ModifyRequestedQuantity(ctx context.Context, input models.ModifyQuantity) (*supply.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateProject(ctx context.Context, input models.CreateProject) (*supply.Project, error) {
	panic("not implemented")
}
func (r *mutationResolver) CloseProject(ctx context.Context, input models.CloseProject) (*supply.Project, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Orders(ctx context.Context, after *string, first *int, before *string, last *int) (*models.OrderConnection, error) {
	return &models.OrderConnection{
		TotalCount: 99,
		PageInfo: &models.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor:     nil,
			EndCursor:       nil,
		},
		Edges: nil,
	}, nil
}
