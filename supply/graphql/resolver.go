//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"errors"
	"field/supply"
	"field/supply/graphql/models"
	"field/supply/ordering"
	"field/supply/search"
	"github.com/99designs/gqlgen/handler"
	"net/http"
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

func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrder) (bool, error) {
	err := r.osvc.CreateOrder(input.OrderID, input.ProjectID)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) SendOrder(ctx context.Context, input models.SendOrder) (bool, error) {
	err := r.osvc.SendOrder(input.OrderID)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) AddOrderItem(ctx context.Context, input models.AddOrderItem) (bool, error) {
	err := r.osvc.AddOrderItem(input.OrderID, input.ProductID, input.Name, input.Uom)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) RemoveOrderItem(ctx context.Context, input models.RemoveOrderItem) (bool, error) {
	err := r.osvc.RemoveOrderItem(input.OrderID, input.ProductID)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) ReceiveOrderItem(ctx context.Context, input models.ModifyQuantity) (bool, error) {
	err := r.osvc.ReceiveOrderItem(input.OrderID, input.ProductID, input.Quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) ModifyRequestedQuantity(ctx context.Context, input models.ModifyQuantity) (bool, error) {
	err := r.osvc.ModifyRequestedQuantity(input.OrderID, input.ProductID, input.Quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Order(ctx context.Context, orderID string) (*supply.Order, error) {
	order, err := r.osvc.FindOrder(orderID)
	if err != nil {
		return &supply.Order{}, errors.New("order not found")
	}
	return order, nil
}
func (r *queryResolver) ProjectOrders(ctx context.Context, projectID string) ([]ordering.ProjectOrder, error) {
	return r.osvc.FindProjectOrderDates(projectID), nil
}
func (r *queryResolver) Products(ctx context.Context, name string) ([]search.Result, error) {
	return r.ssvc.ProductSearch(name), nil
}
